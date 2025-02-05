package linkedlist

import (
	"errors"
)

// DoubleLinkedList یک لیست پیوندی دوطرفه است که هر نود هم به عنصر قبل و هم به عنصر بعد اشاره دارد.
type DoubleLinkedList[T any] struct {
	Head *DoubleNode[T] // اشاره‌گر به اولین نود لیست
	Tail *DoubleNode[T] // اشاره‌گر به آخرین نود لیست
	Size uint           // تعداد نودهای موجود در لیست (چون لیست دو طرفه است در جستجو به ما کمک می‌کند)
}

// DoubleNode یک نود در لیست پیوندی دوطرفه است.
type DoubleNode[T any] struct {
	Data T              // مقدار ذخیره‌شده در نود
	Next *DoubleNode[T] // اشاره‌گر به نود بعدی
	Prev *DoubleNode[T] // اشاره‌گر به نود قبلی
}

// Append مقدار جدیدی را به انتهای لیست پیوندی دوطرفه اضافه می‌کند.
func (l *DoubleLinkedList[T]) Append(data T) {
	newNode := &DoubleNode[T]{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Prev = nil // اگر لیست خالی است، مقدار قبلی نود را nil قرار بده
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail // نود جدید به نود قبلی متصل می‌شود
	}
	l.Tail = newNode // نود جدید، آخرین نود لیست می‌شود
	l.Size++
}

// ItemAt مقدار نودی را که در ایندکس مشخصی قرار دارد، برمی‌گرداند.
func (l *DoubleLinkedList[T]) ItemAt(index uint) (T, error) {
	current := l.Head
	var zeroT T
	if current == nil {
		return zeroT, errors.New("empty list") // اگر لیست خالی باشد، خطا برمی‌گرداند
	}
	length := l.Size
	// بررسی ایندکس نامعتبر
	if index >= length {
		return zeroT, errors.New("index out of bound")
	}
	// انتخاب نقطه شروع جستجو برای بهینه‌سازی
	if index < length/2 {
		// جستجو از ابتدای لیست
		var i uint
		for i = 0; i < index; i++ {
			current = current.Next
		}
	} else {
		// جستجو از انتهای لیست
		current = l.Tail
		for i := length - 1; i > index; i-- {
			current = current.Prev
		}
	}
	return current.Data, nil
}

// Range یک کانال برمی‌گرداند که تمام مقادیر لیست را از ابتدا به انتها ارسال می‌کند.
func (n *DoubleLinkedList[T]) Range() <-chan T {
	ch := make(chan T)
	go func() {
		for current := n.Head; current != nil; current = current.Next {
			ch <- current.Data // مقدار نود فعلی را به کانال ارسال کن
		}
		close(ch) // بعد از ارسال همه داده‌ها، کانال را ببند
	}()
	return ch
}

// RangeBackward یک کانال برمی‌گرداند که تمام مقادیر لیست را از انتها به ابتدا ارسال می‌کند.
func (n *DoubleLinkedList[T]) RangeBackward() <-chan T {
	ch := make(chan T)
	go func() {
		for current := n.Tail; current != nil; current = current.Prev {
			ch <- current.Data // مقدار نود فعلی را به کانال ارسال کن
		}
		close(ch) // بعد از ارسال همه داده‌ها، کانال را ببند
	}()
	return ch
}

func (l *DoubleLinkedList[T]) PopAt(index uint) (T, error) {
	var zeroT T
	if l.Head == nil {
		return zeroT, errors.New("empty list")
	}

	length := l.Size
	// بررسی ایندکس نامعتبر
	if index >= length {
		return zeroT, errors.New("index out of bound")
	}

	var target *DoubleNode[T]

	// انتخاب نقطه شروع جستجو برای بهینه‌سازی
	if index < length/2 {
		// جستجو از ابتدای لیست
		target = l.Head
		var i uint
		for i = 0; i < index; i++ {
			target = target.Next
		}
	} else {
		// جستجو از انتهای لیست
		target = l.Tail
		var i uint
		for i = 0; i > (length-1)-index; i-- {
			target = target.Prev
		}
	}

	// حذف نود از لیست
	if target.Prev != nil {
		target.Prev.Next = target.Next
	} else {
		l.Head = target.Next // حذف اولین نود
	}

	if target.Next != nil {
		target.Next.Prev = target.Prev
	} else {
		l.Tail = target.Prev // حذف آخرین نود
	}

	return target.Data, nil
}

package linkedlist

import (
	"errors"
)

// LinkedList یک اینترفیس کلی برای انواع لیست‌های پیوندی (تک‌جهتی و دوطرفه) است.
// این اینترفیس متدهای پایه‌ای مثل Append، ItemAt و Range را تعریف می‌کند.
type LinkedList[T any] interface {
	Append(data T)               // افزودن یک مقدار جدید به لیست
	ItemAt(index int) (T, error) // دریافت مقدار یک نود بر اساس ایندکس
	Range() <-chan T             // پیمایش عناصر لیست به‌صورت کانال
	PopAt(index int) (T, error)  // حذف یک نود بر اساس ایندکس
}

// SingleLinkedList یک لیست پیوندی تک‌جهتی است که نودهای آن فقط به عنصر بعدی اشاره دارند.
type SingleLinkedList[T any] struct {
	Head *Node[T] // اشاره‌گر به اولین نود لیست
	Tail *Node[T] // اشاره‌گر به آخرین نود لیست
}

// Node یک نود در لیست پیوندی تک‌جهتی است.
type Node[T any] struct {
	Data T        // مقدار ذخیره‌شده در نود
	Next *Node[T] // اشاره‌گر به نود بعدی
}

// Append مقدار جدیدی را به انتهای لیست پیوندی تک‌جهتی اضافه می‌کند.
func (l *SingleLinkedList[T]) Append(data T) {
	newNode := &Node[T]{Data: data}
	if l.Head == nil {
		l.Head = newNode // اگر لیست خالی است، اولین نود را تنظیم کن
	} else {
		l.Tail.Next = newNode // نود جدید را به آخرین نود فعلی متصل کن
	}
	l.Tail = newNode // نود جدید، آخرین نود لیست می‌شود
}

// ItemAt مقدار نودی را که در ایندکس مشخصی قرار دارد، برمی‌گرداند.
func (l *SingleLinkedList[T]) ItemAt(index int) (T, error) {
	current := l.Head
	var zeroT T
	if current == nil {
		return zeroT, errors.New("empty list") // اگر لیست خالی باشد، خطا برمی‌گرداند
	}
	counter := 0
	for current != nil {
		if counter == index {
			return current.Data, nil // مقدار نود در ایندکس مورد نظر را برمی‌گرداند
		}
		counter++
		current = current.Next
	}
	return zeroT, errors.New("index out of bound") // اگر ایندکس معتبر نبود، خطا برمی‌گرداند
}

// Range یک کانال برمی‌گرداند که تمام مقادیر لیست را یکی‌یکی ارسال می‌کند.
// this code has a problem with memory usage
// if outer loop does not read all the values from the channel
// channel will be blocked and goroutine will be leaked
func (n *SingleLinkedList[T]) Range() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch) // بعد از ارسال همه داده‌ها، کانال را ببند
		for current := n.Head; current != nil; current = current.Next {
			ch <- current.Data // مقدار نود فعلی را به کانال ارسال کن
		}
	}()
	return ch
}

func (n *SingleLinkedList[T]) PopAt(index int) (T, error) {
	var zeroT T
	if index == 0 {
		if n.Head == nil {
			return zeroT, errors.New("empty list")
		}
		value := n.Head.Data
		n.Head = n.Head.Next
		return value, nil
	}
	current := n.Head
	counter := 1
	for current != nil {
		if counter == index {
			if current.Next == nil {
				return zeroT, errors.New("index out of bound")
			}
			value := current.Next.Data
			current.Next = current.Next.Next
			return value, nil
		}
		counter++
		current = current.Next
	}
	return zeroT, errors.New("index out of bound")

}

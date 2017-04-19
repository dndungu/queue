package queue

import "testing"

func TestEmpty(t *testing.T) {
	queue := NewQueue()
	if !queue.Empty() {
		t.Errorf("Error, expected Queue to be empty")
	}
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue()
	data := []string{"A", "B", "C", "D", "E", "F"}
	for _, v := range data {
		queue.Enqueue(v)
	}
	first := data[len(data)-1]
	if queue.head.value != first {
		t.Errorf("Error, expected %s, found %s in the head position", first, queue.head.value)
	}
	last := data[0]
	if queue.tail.value != last {
		t.Errorf("Error, expected %s, found %s in the tail position.", last, queue.tail.value)
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{}, []string{"", "", ""}},
		{[]string{"A"}, []string{"A", "", ""}},
		{[]string{"A", "B", "C", "D", "E", "F"}, []string{"A", "B", "C", "D", "E", "F"}},
	}
	for i, test := range tests {
		queue := NewQueue()
		for _, v := range test.input {
			queue.Enqueue(v)
		}
		for _, expected := range test.output {
			actual, exists := queue.Dequeue()
			if exists {
				if actual != expected {
					t.Errorf("%d: Error, expected %s, found %s", i, expected, actual)
				}
			}
		}
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		input    []string
		expected interface{}
	}{
		{[]string{}, nil},
		{[]string{"A"}, "A"},
		{[]string{"A", "B", "C", "D", "E", "F"}, "A"},
	}
	for i, test := range tests {
		queue := NewQueue()
		for _, v := range test.input {
			queue.Enqueue(v)
		}
		actual, exists := queue.Peek()
		if exists {
			if actual != string(test.expected.(string)) {
				t.Errorf("%d: Error, expected %s, found %s", i, test.expected, actual)
			}
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    []string
		expected int64
	}{
		{[]string{}, 0},
		{[]string{"A"}, 1},
		{[]string{"A", "B", "C", "D", "E"}, 5},
	}
	for i, test := range tests {
		queue := NewQueue()
		for _, v := range test.input {
			queue.Enqueue(v)
		}
		actual := queue.Length()
		if actual != test.expected {
			t.Errorf("%d: Error, expected %d, got %d", i, test.expected, actual)
		}
	}
}

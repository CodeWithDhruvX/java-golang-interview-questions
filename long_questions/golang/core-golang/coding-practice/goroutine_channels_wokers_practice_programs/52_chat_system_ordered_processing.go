package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	ID        int
	Content   string
	Sender    string
	Timestamp time.Time
}

type ChatRoom struct {
	messages  chan Message
	processed chan string
	wg        sync.WaitGroup
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		messages:  make(chan Message, 100),
		processed: make(chan string, 100),
	}
}

func (cr *ChatRoom) Start() {
	cr.wg.Add(1)
	go cr.processMessages()
}

func (cr *ChatRoom) processMessages() {
	defer cr.wg.Done()
	defer close(cr.processed)
	
	// Buffer for ordering messages
	messageBuffer := make([]Message, 0)
	
	for msg := range cr.messages {
		messageBuffer = append(messageBuffer, msg)
		
		// Process when buffer is full or no more messages coming
		if len(messageBuffer) >= 5 {
			cr.processBuffer(&messageBuffer)
		}
	}
	
	// Process remaining messages
	cr.processBuffer(&messageBuffer)
}

func (cr *ChatRoom) processBuffer(buffer *[]Message) {
	if len(*buffer) == 0 {
		return
	}
	
	// Sort messages by timestamp
	for i := 0; i < len(*buffer)-1; i++ {
		for j := i + 1; j < len(*buffer); j++ {
			if (*buffer)[i].Timestamp.After((*buffer)[j].Timestamp) {
				(*buffer)[i], (*buffer)[j] = (*buffer)[j], (*buffer)[i]
			}
		}
	}
	
	// Process in order
	for _, msg := range *buffer {
		processed := fmt.Sprintf("[%v] %s: %s", 
			msg.Timestamp.Format("15:04:05.000"), 
			msg.Sender, 
			msg.Content)
		cr.processed <- processed
	}
	
	*buffer = (*buffer)[:0] // Clear buffer
}

func (cr *ChatRoom) SendMessage(msg Message) {
	cr.messages <- msg
}

func (cr *ChatRoom) GetProcessedMessages() <-chan string {
	return cr.processed
}

func (cr *ChatRoom) Stop() {
	close(cr.messages)
	cr.wg.Wait()
}

func sender(name string, room *ChatRoom, wg *sync.WaitGroup) {
	defer wg.Done()
	
	messages := []string{
		"Hello everyone!",
		"How are you?",
		"Great weather today",
		"Anyone up for coffee?",
		"See you later!",
	}
	
	for i, content := range messages {
		msg := Message{
			ID:        i + 1,
			Content:   content,
			Sender:    name,
			Timestamp: time.Now().Add(time.Duration(i) * time.Millisecond),
		}
		
		room.SendMessage(msg)
		fmt.Printf("%s sent: %s\n", name, content)
		time.Sleep(time.Duration(name[0]-'A'+1) * 100 * time.Millisecond)
	}
}

func main() {
	room := NewChatRoom()
	room.Start()
	
	fmt.Println("=== Chat System: Multiple Senders, Ordered Processing ===")
	
	// Start multiple senders
	wg := sync.WaitGroup{}
	senders := []string{"Alice", "Bob", "Charlie"}
	
	for _, sender := range senders {
		wg.Add(1)
		go sender(sender, room, &wg)
	}
	
	// Display processed messages
	go func() {
		for processed := range room.GetProcessedMessages() {
			fmt.Println("Processed:", processed)
		}
	}()
	
	// Wait for all senders
	wg.Wait()
	
	// Stop chat room
	room.Stop()
	
	fmt.Println("Chat system demonstration completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a chat system with multiple senders but ordered message processing?

**Your Response:** I implement a chat room that collects messages from multiple senders and processes them in chronological order. Multiple users send messages concurrently to a central message channel.

The chat room uses a buffering and sorting strategy. It collects messages in a buffer, sorts them by timestamp, then processes them in order. This ensures that even though messages arrive concurrently, they're displayed chronologically.

Each sender creates messages with timestamps and sends them to the chat room. The chat room's processor maintains order by sorting before output, which is crucial for chat applications where message order affects conversation flow.

The key insight is separating concurrent message collection from ordered processing. Senders don't block each other, but the final output maintains proper chronological order for readability.

This pattern is essential for real-time communication systems like chat applications, collaborative editors, or any system where multiple users generate content that needs to be presented in a logical order. It demonstrates understanding of concurrent systems with ordering requirements.

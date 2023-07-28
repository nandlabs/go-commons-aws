# SQS

The SQS library provides a convenient and flexible way to interact with message queues using the Provider interface, which internally contains the Producer and Consumer interfaces. This library allows you to send and receive messages to and from a specific provider, making it easy to work with various URL schemes.

---

* [Introduction](#introduction)
* [Producer Interface](#producer-interface)
* [Receiver Interface](#receiver-interface)
* [Provider Interface](#provider-interface)
* [Installation](#installation)
* [Usage](#usage)
* [Examples](#examples)
* [Contributing](#contributing)

---

### Introduction

The SQS library provides a set of interfaces and functions to work with message queues. It abstracts the underlying implementation details, allowing you to interact with different providers in a unified manner. The central concept of this library is the Provider interface, which combines the functionalities of both the Producer and Consumer interfaces.

### Producer Interface

The Producer interface is used to send messages to a specific provider. It defines the following methods:

* Send(*url.URL, Message, ...Option) error: Sends an individual message to the specified URL.
* SendBatch(*url.URL, []Message, ...Option) error: Sends a batch of messages to the specified URL.

### Receiver Interface

The Receiver interface provides functions for receiving messages from a specific provider. It defines the following methods:

* Receive(*url.URL, ...Option) (Message, error): Performs an on-demand receive of a single message from the specified URL. The behavior of this function may or may not wait for messages to arrive, depending on the implementation.
* ReceiveBatch(*url.URL, ...Option) ([]Message, error): Performs an on-demand receive of a batch of messages from the specified URL. Similarly, the behavior may or may not wait for messages to arrive.
* Additionally, the Receiver interface allows you to register listeners for messages using the AddListener(*url.URL, func(msg Message), ...Option) error method.

### Provider Interface

The Provider interface includes both the Producer and Receiver interfaces, providing a unified interface to interact with message queues. It defines the following methods:

* Producer: Includes the methods from the Producer interface.
* Receiver: Includes the methods from the Receiver interface.
* Schemes() []string: Returns an array of URL schemes supported by the provider.
* Setup(): Performs any necessary setup for the provider.
* NewMessage(string, ...Option) (Message, error): Creates a new message that can be used by clients, expecting the scheme to be provided.

### Installation

To install the SQS library, you can use the following go get command:

```go
go get go.nandlabs.io/commons-aws/sqs
```

### Usage

To use the SQS library in your Go project, import the package and create an instance of the Provider. The library allows you to interact with different message queue providers seamlessly using the same interface.

```go
package main

import (
	"go.nandlabs.io/commons-aws/sqs"
)

func main() {
	// Create a new instance of the SQS provider
	provider := sqs.AddSessionProvider()

	// Use the provider to send and receive messages
	// ...
}

```
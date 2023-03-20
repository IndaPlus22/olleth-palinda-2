# olleth-palinda-2

## Task 2
1. What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
> If you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function, the program may deadlock. This is because the Wait function waits for all producers to finish before closing the channel, but the producers are waiting to send their last messages on the channel, and the consumers are waiting to receive those messages. If the channel is not closed before the producers and consumers finish, the program will hang.

> If you move the close(ch) from the main function and instead close the channel in the end of the function Produce, the program will still work correctly as long as all producers have finished before closing the channel. This is because closing the channel signals to the consumers that there will be no more messages, and they can exit their loops. If a producer tries to send a message on a closed channel, it will panic.

2. What happens if you remove the statement close(ch) completely?
> If you remove the statement close(ch) completely, the program may deadlock. This is because the consumers are waiting to receive messages from the channel, and if the producers have finished sending messages but the channel is not closed, the consumers will block forever waiting for more messages.
3. What happens if you increase the number of consumers from 2 to 4?
> If you increase the number of consumers from 2 to 4, the program will still work correctly. The consumers will share the work of receiving messages from the channel, and they will print the messages they receive as soon as they receive them.
4. Can you be sure that all strings are printed before the program stops?
> It is guaranteed that all strings will be printed before the program stops, as long as the channel is closed after all producers have finished sending messages. This is because the consumers will keep receiving messages until the channel is closed, and they will print each message as soon as they receive it. Once the channel is closed, the consumers will exit their loops and the program will stop.
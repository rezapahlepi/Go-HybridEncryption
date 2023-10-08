# Go-HybridEncryption
Hybrid encryption is an encryption technique that combines two types of encryption: symmetric encryption and asymmetric encryption. This is used to overcome several challenges in data encryption that exist in asymmetric encryption and symmetric encryption.

In hybrid encryption, the message to be sent is encrypted using symmetric encryption with a random key called a "session key." This session key is only used for one message exchange.
Then, the session key is encrypted using asymmetric encryption (usually RSA or ECC) with the recipient's public key.
The encrypted message and the encrypted session key are sent to the recipient.
The recipient uses their private key to decrypt the session key, and then they use the session key to decrypt the symmetric message.

Advantages of Hybrid Encryption:

Security: Randomly generated session keys are used for one exchange only, reducing the risk of repeated use of the same key.
Speed: Most encryption and decryption is done with faster symmetric encryption, saving time and computing resources.
Key Management: Public keys can be freely distributed, eliminating the need to overcome complex public key management issues. The private key is only needed by the recipient of the message.
Flexibility: Hybrid encryption makes it possible to combine the advantages of both types of encryption (symmetric and asymmetric) to provide a high level of security and efficiency.

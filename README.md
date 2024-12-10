# Wallet Word Scrambler

This program helps you securely transform your existing wallet words (from a SLIP39 wordlist) into a new set of wallet words using a password and optional salt words. It is designed for use in an **air-gapped environment** to enhance security and prevent unauthorized access to sensitive information.

## Features

- **Uses SLIP39 Wordlist**: A standard 1024-word list for wallet backups.
- **Password Protection**: Derives cryptographic keys using a password you provide.
- **Salt Support**: Allows you to provide additional entropy with manually entered or randomly generated salt words.
- **Secure Key Derivation**: Utilizes Argon2 and SHA3-256 for cryptographic operations.
- **Air-Gapped Usage**: Designed to run on a machine disconnected from any network for maximum security.

---

## Requirements

- **Go** (Golang) installed (version 1.16+ recommended).

---

## Usage Instructions

### 1. **Prepare an Air-Gapped Environment**
   - Format a machine and ensure it has no network connectivity.
   - Install Go if not already installed.

### 2. **Run the Program**
   - Compile the program using Go:
     ```bash
     go build -o wallet-scrambler .
     ```
   - Run the compiled executable:
     ```bash
     ./wallet-scrambler
     ```

### 3. **Follow the Prompts**
   - **Password Setup**:
     - Enter a password twice to confirm. Ensure you remember it as there is **no recovery option**.
     - Weak passwords will prompt a warning, but you can choose to proceed.
   - **Salt Words**:
     - You can either:
       - Enter salt words manually (`manual` mode).
       - Generate random salt words (`random` mode).
     - Salt words enhance the security of the key derivation process.
   - **Wallet Words**:
     - Input the number of words in your wallet (4–33).
     - Enter each wallet word when prompted. Each word must exist in the SLIP39 wordlist.
   - The program will calculate a new set of wallet words using the provided password, salt, and input wallet words.

---

## Output

The program will display two sets of words:

1. **Salt Words**: The salt words you provided or generated.
2. **New Wallet Words**: A new set of words derived from your input.

---

## Security Warning

1. **Air-Gapped Environment**: This program should only be run on a **formatted and air-gapped machine**.
   - Avoid running it on a machine connected to any network.
   - Wipe the machine after use to ensure no residual sensitive data remains.
2. **Password Safety**:
   - Your password is crucial for recovering the scrambled wallet words.
   - **Never forget your password**, as there is no way to recover it.
3. **Storage**:
   - Write down the generated **salt words** and **new wallet words** and store them securely.

---

## Example Session

**Input**:
- Password: `MySecureP@ssw0rd`
- Salt: `manual` (entered: `apple`, `banana`, `cherry`, `date`)
- Wallet Words: (entered: `academic`, `biology`, `energy`, `gravity`)

**Output**:
```plaintext
*** Here are your new wallet words ***

Salt:
1. apple
2. banana
3. cherry
4. date

Wallet Words:
1. express
2. physics
3. valid
4. guitar

Write both salt and words down and store them in a safe place.

*** NEVER forget your password - there is no way to recover it ***
```

---

## Notes

- **Wordlist**: The SLIP39 English wordlist is embedded in the program.
- **Performance**: Key derivation is intentionally slow for security reasons.

---

Let me know if you'd like additional details or modifications!

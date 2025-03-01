# Password Manager

This is a simple cli password manager written in Go. It allows users to create accounts, log in, and manage their passwords securely. Each user's passwords are stored in a separate encrypted file, ensuring that only the respective user can access their passwords.

## Features

- Create an account
- Log in to an account
- Log out from an account
- Add a password for a service
- Generate a random password
- List all stored passwords
- Clear the terminal
- Exit the program

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/Armir2005/password-manager.git
   cd password-manager
   ```

2. **Build the program:**

    ```sh
    go build -o password-manager
    ```

3. **Run the program:**

    ```sh
    ./password-manager
    ```

## Usage

When you run the program, you will see a prompt where you can enter commands. Here are the available commands:

- **help:** Show available commands.
- **create [Username] [Password]:** Create an account with the specified username and password.
- **login [Username] [Password]:** Log in to your account with the specified username and password.
- **logout:** Log out from your account.
- **add [Service] [Username] [Password]:** Add a password for the specified service and username.
- **generate [Service] [Username] [Length] [Special Characters (y/n)]:** Generate a random password for the specified service and username. Specify the length and whether to include special characters.
- **list:** List all stored passwords.
- **clear:** Clear the terminal.
- **exit:** Exit the program.

## Security

- Passwords are encrypted using AES-GCM with a key derived from the user's password and a unique salt.
- Each user's passwords are stored in a separate file named **[username]_passwords.json**.
- The salt is generated randomly for each user and stored alongside the encrypted password.

## Example

    ```sh
    Password Manager:
    Type 'help', to show available commands.
    > create user1 password1
    Account saved.
    > login user1 password1
    Login successful.
    > add github user1 mypassword
    Password saved.
    > list
    Service: github, Username: user1, Password: mypassword
    > logout
    Logged out successfully.
    > exit
    Exiting...
    ```

## Licence

This project is licensed under the MIT License.
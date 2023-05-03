# NITH-Online-Internship-Document-Signing

Welcome to the app! To get started, you'll need to follow these steps to set up the backend on your local machine.

## Prerequisites

Before you get started, you'll need to make sure you have the following installed on your machine:
e
- GoLAng (version 20.2 or later)

## Installation

1. Clone the repository from GitHub:

        git clone https://github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing

2. Navigate into the project directory:

        cd back-end

3. Install the dependencies:
        go get .
        go mod tidy
        
5. Setup .env file:
        
        path: ./back-end
        
        Format:
        ```.env
        PORT=:x

        DATABASE_URL=[postgress]

        REDIS_URL=

        ACCESS_TOKEN_PRIVATE_KEY=
        ACCESS_TOKEN_PUBLIC_KEY=
        ACCESS_TOKEN_EXPIRED_IN=15m
        ACCESS_TOKEN_MAXAGE=15

        REFRESH_TOKEN_PRIVATE_KEY=
        REFRESH_TOKEN_PUBLIC_KEY=
        REFRESH_TOKEN_EXPIRED_IN=60m
        REFRESH_TOKEN_MAXAGE=60
        

4. Start the development server:

        go run dev/go run main.go 

That's it! You can now have the app running locally on your machine.

TODO: Use uber fx!
more routes needed! and better file structure.

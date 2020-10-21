# grpc/grpc issue #24469

To trigger the error add a broken cert to the `ssl` folder and run `make` to start the server.
Once it's running you can use either of the options below to trigger the client.

# Setting up the client

You should open `client/index.php` and change the `$caPath`-variable to where your CA's are stored.

# Start the client

## Plain
```bash
php ./client/index.php
```

## PHP server
```bash
cd client && php -S localhost:8000
```

Visiting http://localhost:8000/ will crash the server.

## PHP-FPM using `symfony serve`
```bash
cd client && symfony serve --no-tls --dir=.
```

Visiting http://localhost:8000/ you'll be greeted by an error.

> **Note**: For more information about the `symfony`-cli and how to install it, checkout [Symfony.com](https://symfony.com/doc/current/setup/symfony_server.html#installation)

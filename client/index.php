<?php
require __DIR__.'/vendor/autoload.php';

$caPath = '/usr/local/etc/openssl@1.1/cert.pem';

$client = new Helloworld\GreeterClient(
    'localhost:50051',
    ['credentials' => Grpc\ChannelCredentials::createSsl(file_get_contents($caPath))]
);

$request = new Helloworld\HelloRequest();
$request->setName("stanley-cheung");

list($response, $status) = $client->SayHello($request)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: ".$status->code.", ".$status->details.PHP_EOL;
    exit($status->code);
}

echo $response->getMessage().PHP_EOL;

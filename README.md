# Garage

## Overview

Garage is an application that allows users to book visits in the nearest car repair shops. Project carried out as an engineering thesis at the Silesian University of Technology.

## Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Node.js](https://nodejs.org/en/download/package-manager)
- [Expo CLI](https://www.npmjs.com/package/expo-cli)

## Installation

1. Clone the Repository.
```bash
git clone https://github.com/KsaweryZietara/garage.git
cd garage
```

2. Run tests.
```bash
make test
```

3. Set environment variables defined in [`compose.yaml`](https://github.com/KsaweryZietara/garage/blob/main/compose.yaml).

4. Start the application.
```bash
make run
```

5. Launch the user interface. 
```bash
make web
```

User interface is available on port 8081. You can access it by navigating to:
```
http://localhost:8081
```

> [!NOTE]
> To use application on mobile device download [expo](https://play.google.com/store/apps/details?id=host.exp.exponent&referrer=docs) and replace `localhost` with server IP address.

6. Stop the application.
```bash
make stop
```

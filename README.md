# status-app-web
Reimplementating the Web App of (https://mysms4you.com/) using Go, Firestore, Firebase and yet to come. 

A few resources to get you started

- [GoLang](https://golang.org/)
- [Firestore](https://firebase.google.com/products/firestore)
- [Authentication](https://firebase.google.com/products/auth)
- [Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)
- [Cloud Storage](https://cloud.google.com/storage)
- [App Engine](https://cloud.google.com/appengine)
- [Firebase Admin SDK](https://firebase.google.com/docs/reference/admin)
- [Firebase Emulator Suit](https://firebase.google.com/docs/emulator-suite)

## Build & Run

* Change `ProjectID` in `data_constant.go`
* If using Firebase emulator 
  - Change `Emulator_Port` in `data_constant.go`

#### Execute inside Project directory & Follow the instruction

```console
firebase init
firebase emulators:start

go build src/main/main.go
go run src/main/main.go
```

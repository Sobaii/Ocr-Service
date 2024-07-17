# Ocr-Service
Optical character recognition (OCR) file processing service

### Start server
Ensure `127.0.0.1:50051` and `localhost:50052` are not in use. Run `air` in base directory.

### Updating stubs

#### Generate golang code
Run `protoc -I . -I proto/googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=. proto/ocr_service.proto`

#### Generate typescript code (for web client)
Run `protoc -I=proto ocr_service.proto   --js_out=import_style=commonjs:./client-stubs   --grpc-web_out=import_style=typescript,mode=grpcwebtext:./client-stubs`
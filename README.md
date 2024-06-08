# Ocr-Service
Optical character recognition (OCR) file processing service

### Updating stubs

#### Generate golang code
run `protoc --go_out=paths=source_relative:./ocr --go-grpc_out=paths=source_relative:./ocr proto/ocr_service.proto`

#### Generate typescript code (for web client)
run `protoc -I=proto ocr_service.proto   --go_out=paths=source_relative:./ocr   --go-grpc_out=paths=source_relative:./ocr   --js_out=import_style=commonjs:./client-stubs   --grpc-web_out=import_style=typescript,mode=grpcwebtext:./client-stubs`
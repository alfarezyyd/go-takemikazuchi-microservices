package dto

import "mime/multipart"

type CreateWorkerWalletDocumentDto struct {
	IdentityCard      *multipart.FileHeader
	PoliceCertificate *multipart.FileHeader
	DriverLicense     *multipart.FileHeader
}

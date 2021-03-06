/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// BuildSignatureKeyType : Public key formats.   - KEY_TYPE_UNSPECIFIED: `KeyType` is not set.  - PGP_ASCII_ARMORED: `PGP ASCII Armored` public key.  - PKIX_PEM: `PKIX PEM` public key.
type BuildSignatureKeyType string

// List of BuildSignatureKeyType
const (
	KEY_TYPE_UNSPECIFIED_BuildSignatureKeyType BuildSignatureKeyType = "KEY_TYPE_UNSPECIFIED"
	PGP_ASCII_ARMORED_BuildSignatureKeyType BuildSignatureKeyType = "PGP_ASCII_ARMORED"
	PKIX_PEM_BuildSignatureKeyType BuildSignatureKeyType = "PKIX_PEM"
)

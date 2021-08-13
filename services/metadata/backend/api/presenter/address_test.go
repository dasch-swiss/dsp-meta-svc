package presenter_test

import (
	"testing"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/api/presenter"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
)

func Test_AddressPresenter_NullifyJsonProps(t *testing.T) {
	const ADDRESS_TYPE = "http://ns.dasch.swiss/repository#Address"
	id, _ := valueobject.NewIdentifier()

	a := presenter.Address{
		ID:         id,
		Type:       ADDRESS_TYPE,
		Street:     "street",
		PostalCode: "0000",
		Locality:   "city",
		Country:    "country",
		Canton:     "canton",
		Additional: "additional",
		CreatedAt:  "2021-08-05 12:12:00 +0000 UTC",
		CreatedBy:  "12345678-1234-1234-1234-123456789101",
		ChangedAt:  "0001-01-01 00:00:00 +0000 UTC",
		ChangedBy:  "00000000-0000-0000-0000-000000000000",
		DeletedAt:  "0001-01-01 00:00:00 +0000 UTC",
		DeletedBy:  "00000000-0000-0000-0000-000000000000",
	}

	a = a.NullifyJsonProps()

	assert.Equal(t, a.ID, id)
	assert.Equal(t, a.Type, ADDRESS_TYPE)
	assert.Equal(t, a.Street, "street")
	assert.Equal(t, a.PostalCode, "0000")
	assert.Equal(t, a.Locality, "city")
	assert.Equal(t, a.Country, "country")
	assert.Equal(t, a.Canton, "canton")
	assert.Equal(t, a.Additional, "additional")
	assert.Equal(t, a.CreatedAt, "2021-08-05 12:12:00 +0000 UTC")
	assert.Equal(t, a.CreatedBy, "12345678-1234-1234-1234-123456789101")
	assert.Equal(t, a.ChangedAt, "")
	assert.Equal(t, a.ChangedBy, "")
	assert.Equal(t, a.DeletedAt, "")
	assert.Equal(t, a.DeletedBy, "")
}

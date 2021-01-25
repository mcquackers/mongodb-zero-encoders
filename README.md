This package is to remedy the official MongoDB Golang driver's unconventional support of the `omitempty` struct flag.

In most encoders, the `omitempty` flag omits zero values from being marshalled to the outputted bytes.  The MongoDB driver's
behavior continues to encode zero value primitives into the outputted bytes.  To rectify this, this package creates structs
for primitive encoders, farming out the actual encoding to bsoncodec.DefaultValueEncoders, or specific encoders where needed, but also adds methods to implement
the bsoncodec.TypeZeroer interface, allowing zero value primitives to be successfully omitted from marshalled bytes.

Usage:

For default encoders + zero value omitting encoders:
////////////////
import(
	zero_enc  "github.com/mcquackers/mongodb-zero-encoders"
	"go.mongodb.org/mongo-driver/mongo/options"
)
reg := zero_enc.DefaultValueEncoders()
clientOptions := options.ClientOptions().SetRegistry(reg)
////////////////

For registering these encoders on an existing Registry:

////////////////
import(
        zero_enc  "github.com/mcquackers/mongodb-zero-encoders"
        "go.mongodb.org/mongo-driver/bson/bsoncodec""
)
rb := bsoncodec.NewRegistryBuilder()
//add your own encoders to registry
zero_enc.RegisterZeroEncoders(rb)
reg := rb.Build()
clientOptions := options.ClientOptions().SetRegistry(reg)
////////////////


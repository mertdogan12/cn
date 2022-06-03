set -e

mongo <<EOF

use $MONGO_DATABASE

db.computers.insertOne({
  id: "123412123",
  name: "Moin moin"
})

EOF

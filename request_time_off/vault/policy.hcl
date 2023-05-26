# This section grants access to "kv-v2/data/api-key"
path "kv-v2/data/api-key" {
  capabilities = ["read", "update", "list","create", "delete"]
}


# Mount the AppRole auth method
path "sys/auth/approle" {
  capabilities = [ "create", "read", "update", "delete", "sudo" ]
}

# Configure the AppRole auth method
path "sys/auth/approle/*" {
  capabilities = [ "create", "read", "update", "delete" ]
}

# Create and manage roles
path "auth/approle/*" {
  capabilities = [ "create", "read", "update", "delete", "list" ]
}

# Write ACL policies
path "sys/policies/acl/*" {
  capabilities = [ "create", "read", "update", "delete", "list" ]
}

# Write test data
# Set the path to "secret/data/postgres/*" if you are running `kv-v2`
path "secret/data/postgres/*" {
  capabilities = [ "create", "read", "update", "delete", "list" ]
}

# by Vault to handle generation of dynamic database credentials.
path "database/creds/request_time_off" {
  capabilities = ["read","create", "delete", "list", "update"]
}
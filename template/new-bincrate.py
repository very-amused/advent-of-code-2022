import tomlkit
import sys
import os

if len(sys.argv) != 2:
  raise UserWarning("Missing/invalid binary crate name, exiting")
  os.exit(1)
crate = sys.argv[1]

with open('Cargo.toml') as f:
  toml = tomlkit.load(f)
  members = toml['workspace']['members']
  if crate not in members:
    members.append(crate)
    with open('Cargo.toml', 'w') as f:
      tomlkit.dump(toml, f)
    print(f'Added crate {crate} to Cargo.toml')
  else:
    print(f'Crate {crate} already present in Cargo.toml')

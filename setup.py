import os

# Run a shell command
print("running: go mod vendor")
os.system("go mod vendor")
print("done")

print("running: go mod tidy")
os.system("go mod tidy")
print("done")


user = input("What is the database username (root used in example): ")
password = input("What is the database password: ")
print("If the provided username or password is incorrect, the database creation will fail and application won't connect to db.")
database = input("What is the database name: ")
host = input("What is the database host (localhost used in example): ")
port = input("What is the database port (3306 used in example): ")
jwt_secret = input("Select a secret for JWT cookies: ")
salt_rounds = input("Select a number of salt rounds for password hashing (10 default): ")

lines = []
lines.append(f'DB_USERNAME="{user}"')
lines.append(f'DB_PASSWORD="{password}"')
lines.append(f'DB_HOST="{host}"')
lines.append(f'DB_PORT="{port}"')
lines.append(f'DB_NAME="{database}"')
lines.append(f'salt_rounds={salt_rounds}')
lines.append(f'jwt_secret="{jwt_secret}"')

print("Writing to .env file")
with open(".env", "w") as file:
    file.write("\n".join(lines))
print("done")

print(f"Creating database with name {database}")
os.system(f"mysql -u {user} -p{password} -e 'CREATE DATABASE {database}'")
print(f"Database {database} created, copying dump.sql to it now")
os.system(f"mysql -u {user} -p{password} {database} < dump.sql")
print("done")
print("\n\nSetup complete! Run 'go run main.go' to start the server.")
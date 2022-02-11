module myapp

go 1.17

replace github.com/RedHoodJT1988/velocidad => ../velocidad

require (
	github.com/RedHoodJT1988/velocidad v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
)

require github.com/joho/godotenv v1.4.0 // indirect

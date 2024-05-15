package common

// Error due to unique constraint violation
const PgErrorUniqueViolation = "23505"

// Error due to database shutdown (caused by administrator)
const PgErrorAdminShutdown = "57P01"

// Error due to database shutdown (crash)
const PgErrorCrashShutdown = "57P02"

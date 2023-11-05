-- CreateTable
CREATE TABLE "user" (
    "id" TEXT NOT NULL,
    "password_digest" BYTEA NOT NULL,
    "name" TEXT
);

-- CreateIndex
CREATE UNIQUE INDEX "user_id_key" ON "user"("id");

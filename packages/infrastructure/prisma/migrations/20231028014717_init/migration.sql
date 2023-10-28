-- CreateTable
CREATE TABLE "owner" (
    "id" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "circle" (
    "id" TEXT NOT NULL,
    "owner_id" TEXT NOT NULL,
    "path_digest" BYTEA NOT NULL,
    "path" BYTEA NOT NULL,
    "path_iv" BYTEA NOT NULL,
    "name" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "event" (
    "id" TEXT NOT NULL,
    "circle_id" TEXT NOT NULL,
    "path_digest" BYTEA NOT NULL,
    "path" BYTEA NOT NULL,
    "path_iv" BYTEA NOT NULL,
    "name" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "owner_id_key" ON "owner"("id");

-- CreateIndex
CREATE UNIQUE INDEX "circle_id_key" ON "circle"("id");

-- CreateIndex
CREATE UNIQUE INDEX "circle_path_digest_key" ON "circle"("path_digest");

-- CreateIndex
CREATE UNIQUE INDEX "circle_owner_id_name_key" ON "circle"("owner_id", "name");

-- CreateIndex
CREATE UNIQUE INDEX "event_id_key" ON "event"("id");

-- CreateIndex
CREATE UNIQUE INDEX "event_path_digest_key" ON "event"("path_digest");

-- AddForeignKey
ALTER TABLE "circle" ADD CONSTRAINT "circle_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "owner"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "event" ADD CONSTRAINT "event_circle_id_fkey" FOREIGN KEY ("circle_id") REFERENCES "circle"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- CreateTable
CREATE TABLE "user" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "event" (
    "id" SERIAL NOT NULL,
    "uuid" TEXT NOT NULL,
    "organizer_id" INTEGER NOT NULL,
    "path_digest" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "event_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "event_uuid_key" ON "event"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "event_path_digest_key" ON "event"("path_digest");

-- AddForeignKey
ALTER TABLE "event" ADD CONSTRAINT "event_organizer_id_fkey" FOREIGN KEY ("organizer_id") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- CreateTable
CREATE TABLE "user" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "circle" (
    "id" SERIAL NOT NULL,
    "uuid" TEXT NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "path_digest" TEXT NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "circle_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "event" (
    "id" SERIAL NOT NULL,
    "uuid" TEXT NOT NULL,
    "circle_id" INTEGER NOT NULL,
    "path_digest" TEXT NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "event_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "circle_uuid_key" ON "circle"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "circle_path_digest_key" ON "circle"("path_digest");

-- CreateIndex
CREATE UNIQUE INDEX "event_uuid_key" ON "event"("uuid");

-- CreateIndex
CREATE UNIQUE INDEX "event_path_digest_key" ON "event"("path_digest");

-- AddForeignKey
ALTER TABLE "circle" ADD CONSTRAINT "circle_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "event" ADD CONSTRAINT "event_circle_id_fkey" FOREIGN KEY ("circle_id") REFERENCES "circle"("id") ON DELETE CASCADE ON UPDATE CASCADE;

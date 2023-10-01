-- CreateTable
CREATE TABLE "owner" (
    "id" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "circle" (
    "id" TEXT NOT NULL,
    "owner_id" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "name" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "event" (
    "id" TEXT NOT NULL,
    "circle_id" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "name" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "owner_id_key" ON "owner"("id");

-- CreateIndex
CREATE UNIQUE INDEX "circle_id_key" ON "circle"("id");

-- CreateIndex
CREATE UNIQUE INDEX "circle_path_key" ON "circle"("path");

-- CreateIndex
CREATE UNIQUE INDEX "event_id_key" ON "event"("id");

-- CreateIndex
CREATE UNIQUE INDEX "event_path_key" ON "event"("path");

-- AddForeignKey
ALTER TABLE "circle" ADD CONSTRAINT "circle_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "owner"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "event" ADD CONSTRAINT "event_circle_id_fkey" FOREIGN KEY ("circle_id") REFERENCES "circle"("id") ON DELETE CASCADE ON UPDATE CASCADE;

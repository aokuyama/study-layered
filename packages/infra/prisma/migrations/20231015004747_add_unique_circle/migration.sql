/*
  Warnings:

  - A unique constraint covering the columns `[owner_id,name]` on the table `circle` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "circle_owner_id_name_key" ON "circle"("owner_id", "name");

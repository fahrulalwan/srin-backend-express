const express = require("express");
const router = express.Router();

/**
 * @openapi
 * /brands:
 *   get:
 *     tags:
 *      - brands
 *     description: Get Brand API
 *     responses:
 *       200:
 *         description: Returns a mysterious string.
 */
router.get("/", (req, res, next) => {
  res.send("Get Brand API");
});

/**
 * @openapi
 * /brands:
 *   post:
 *     tags:
 *      - brands
 *     description: Post Brand API
 *     responses:
 *       200:
 *         description: Returns a mysterious string.
 */
router.post("/", (req, res, next) => {
  res.send("Post Brand API");
});

/**
 * @openapi
 * /brands:
 *   put:
 *     tags:
 *      - brands
 *     description: Put Brand API
 *     responses:
 *       200:
 *         description: Returns a mysterious string.
 */
router.put("/", (req, res, next) => {
  res.send("Put Brand API");
});

/**
 * @openapi
 * /brands:
 *   delete:
 *    tags:
 *      - brands
 *    description: Delete Brand API
 *    responses:
 *       200:
 *         description: Returns a mysterious string.
 */
router.delete("/", (req, res, next) => {
  res.send("Delete Brand API");
});

module.exports = router;

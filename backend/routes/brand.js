const express = require('express');
const router = express.Router();


/* GET brand. */
router.get('/', (req, res, next) => {
  res.send('Hello World')
});

/* POST brand. */
router.post('/', (req, res, next) => {
  res.send(req.body)
});

/* PUT brand. */
router.put('/', (req, res, next) => {
  res.send('Hello World')
});

/* DELETE brand. */
router.delete('/', (req, res, next) => {
  res.send('Hello World')
});

module.exports = router;

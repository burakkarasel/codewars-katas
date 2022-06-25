function isValidWalk(walk) {
    const obj = {"n": 0, "s": 0, "e": 0, "w": 0}
    walk.forEach(item => obj[item]++)
    return obj['n'] === obj['s'] && obj['w'] === obj['e'] && walk.length === 10
  }
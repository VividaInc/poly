package cache

import "math/rand"

func (s *CacheSession) collect_hash_pool() []string {
  var buf []string
  for _, hash := range s.Caches {
    buf = append(buf, hash.(string))
  }
  return buf
}

func (s *CacheSession) generate_hash(length int, pool []string) string {
  var (
    chars []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
    buf   []rune
  )
  for i := 0; i < length; i++ {
    if !s.validate_hash(string(buf), pool) {
      s.generate_hash(length, pool)
    }
    buf = append(buf, chars[rand.Intn(len(chars))])
  }
  return string(buf)
}

func (s *CacheSession) validate_hash(buf string, pool []string) bool {
  valid := true
  for _, hash := range pool {
    if buf == hash {
      valid = false
    }
  }
  return valid
}

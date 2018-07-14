package cache

import "github.com/VividaInc/poly/file"

type Cache interface {
  NewSession()
  Save()
  Delete()
  Update()
}

type CacheSession struct {
  Caches map[string]interface{}
}

func NewSession() *CacheSession {
  return &CacheSession{map[string]interface{}{}}
}

func (s *CacheSession) Delete() error {
  return nil
}

func (s *CacheSession) Save(name string, dir string) error {
  pre := s.parse_file_prefix(name)
  ext := file.Parse(name)
  s.Caches[pre] = s.concat_pre_w_ext(s.generate_hash(16,
    s.collect_hash_pool()), ext)
  if err := s.transfer_file_to_store(pre, dir, ext); err != nil {
    return err
  }
  return nil
}

func (s *CacheSession) Update() error {
  return nil
}

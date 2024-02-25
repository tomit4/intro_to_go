## Mutations

### Insert And Element

```go
m[key] = elem
```

### Get An Element

```go
elem = m[key]
```

### Delete An Element

```go
delete(m, key)
```

### Check If A Key Exists

```go
elem, ok := m[key]
```

If `key` is in `m`, then `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element
type.

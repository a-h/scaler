## Scaler

Scale input ranges to output ranges.

```go
inputMinimum := 0
inputMaximum := 10

outputMinimum := 0
outputMaximum := 100

s := scaler.New(inputMinimum, inputMaximum, outputMinimum, outputMaximum)
a, err := s.Scale(0) // 0
b, err := s.Scale(5) // 50
c, err := s.Scale(10) // 100
d, err := s.Invert(100) // 10
e, err := s.Invert(50) // 5
f, err := s.Invert(0) // 0
```
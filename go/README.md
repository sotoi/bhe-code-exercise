# BHE Software Engineer Coding Exercise

## Implementation Research

### Sieve Algorithm Selection
The implementation uses an optimized version of the Sieve of Eratosthenes algorithm only odd numbers being taken into consideration. After researching some prime number generation methods, including:

- Basic Sieve of Eratosthenes
- Sieve of Atkin
- Segmented Sieve

The Odd number Sieve of Eratosthenes was chosen for its:
- Simple implementation
- Good performance characteristics
- Memory efficiency potential

### Memory Optimization
Implemented solution using odd-number-only sieve, which provides several advantages:

1. **Reduced Memory Usage**:
   - Only stores odd numbers in the sieve array
   - Cuts memory requirements in half compared to regular sieve

2. **Performance Benefits**:
   - Fewer iterations needed while sieving
   - Better cache utilization due to smaller memory footprint
   - No need to check even numbers

## Running the Service

Start the server using:
```bash
go run api/main.go
```

The service will start on `http://localhost:8080`


## API Endpoints

### Get Nth Prime Number

```http
GET /primes?n={number}
```

| Parameter | Type    | Description                    |
|-----------|---------|--------------------------------|
| n         | integer | Position of prime to retrieve  |

#### Example Request

Using curl:
```bash
curl "http://localhost:8080/primes?n=0"
```

#### Success Response

```json
{
    "nth_prime": 2
}
```

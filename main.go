package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Sum the distinct multiples of a and b that are less than n.
// Project Euler prob 1.
// Implementation is in linear time (0(n)).
func sumDistinctMultsLinear(a, b, ub int) int {
	sum := 0
	for i := a; i <= ub; i += a {
		sum += i
	}
	for i := b; i <= ub; i += b {
		sum += i
	}
	for i := a * b; i <= ub; i += a * b {
		sum -= i
	}
	return sum
}

// Sum of all multiples of d less or equal to ub.
func sumSeqConst(d int, ub int) int {
	n := ub / d
	ret := n * ((2 * d) + ((n - 1) * d)) / 2
	return ret

}

// Sum the distinct multiples of a and b that are less than n.
// Project Euler prob 1.
// Implementation is in constant time (0(1)).
func sumDistinctMultsConst(a, b, n int) int {
	return sumSeqConst(a, n) + sumSeqConst(b, n) - sumSeqConst(a*b, n)
}

// Sum the even Fibonacci numbers that don't exceed max.
// Project Euler prob 2.
func sumEvenFibs(max int) int {
	prevprev, prev := 1, 2

	sum := prev
	for {
		cur := prevprev + prev
		if cur > max {
			break
		}
		if cur%2 == 0 {
			sum += cur
		}
		prevprev, prev = prev, cur
	}
	return sum
}

// Sieve of Eratosthenes (O(log n)) to find primes up to n.
func sieve(n int) *[]bool {
	notPrimeMask := make([]bool, n+1) // elements default to false
	for i := 2; i <= n; i++ {
		if notPrimeMask[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			notPrimeMask[j] = i != j
		}
	}
	return &notPrimeMask
}

// Biggest prime factor of n.
// Proj Euler prob 3.
func biggestPrimeFactor(n int) int {
	var bpf int

	floorSqrt := int(math.Floor(math.Sqrt(float64(n))))
	for val, isNotPrime := range *sieve(floorSqrt) {
		if val < 2 || isNotPrime {
			continue
		}
		// Do we have a prime factor?
		if n%val == 0 {
			// Yes
			n /= val
			bpf = val
			if n == 1 {
				// Fully factorised
				break
			}
		}
	}

	return bpf
}

// Whether x is a palindrome.
func isPalindrome(x int) bool {
	ret := true
	xRunes := []rune(strconv.Itoa(x))
	lenXRunes := len(xRunes)
	for i := range xRunes {
		if xRunes[i] != xRunes[lenXRunes-i-1] {
			ret = false
			break
		}
	}
	return ret
}

// The largest palindrome made from the product of two n-digit numbers.
// Returns a non-nil error if no palindrome found.
// Project Euler prob 4.
func maxPalinProduct2NDigitInts(nDigits int) (int, error) {
	ub := int(math.Pow10(nDigits)) - 1
	lb := int(math.Pow10(nDigits - 1))
	max := 0
	for i := ub; i >= lb; i-- {
		for j := i; j >= lb; j-- {
			p := i * j
			if isPalindrome(p) {
				if p > max {
					max = p
				}
			}
		}
	}
	var err error
	if max == 0 {
		err = errors.New("No palindromes found")
	}
	return max, err
}

// Greatest Common Divisor of a and b.
// Implemented using Euclid's algorithm.
func gcd(a, b int) int {
	var ret int
	x, y := a, b
	if b < a {
		x, y = b, a
	}
	for {
		rem := y % x // 171
		if rem == 0 {
			ret = x
			break
		}
		x, y = rem, x
	}
	return ret
}

// Absolute value of x.
func absInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Lowest Common Multiple of a and b.
// Implemented using gcd (Greatest Common Multiple) function.
func lcm(a, b int) int {
	return absInt(a*b) / gcd(a, b)
}

// Lowest Common Multiple of all natural numbers up to and including n.
// Proj Euler prob 5.
func lcmRange(n int) int {
	if n < 3 {
		return n
	}
	accum := 2
	for i := 3; i <= n; i++ {
		accum = lcm(accum, i)
	}
	return accum
}

// a to the power of b.
func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

// Difference between the sum of the squares of the first n natural numbers and the square of the sum of those natural numbers.
// Proj Euler prob 6.
func diffSumSqsSqSums(n int) int {
	sumseq := sumSeqConst(1, n)
	// The sums of squares of sequence of integers are 'pyramid' numbers,
	// for which there a 0(1) way of calculating the nth:
	sumSqs := ((2 * n * n * n) + (3 * n * n) + n) / 6
	return (sumseq * sumseq) - sumSqs
}

// Nth prime number.
// Find by iteratively apply Sieve of Eratosthenes over increasingly larger ranges.
// Proj Euler prob 7.
func nthPrime(n int) int {
	ub := 2 * n // upper bound of sieving
	for {
		nPrimes := 0
		for val, isNotPrime := range *sieve(ub) {
			if val >= 2 && !isNotPrime {
				nPrimes++
				if nPrimes == n {
					return val
				}
			}
		}
		ub += n
	}
}

var prob8Runes = []rune(
	"73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450")

// Find largest product of a window of n digits in prob8runes.
func prob8(winSize int) int {
	maxProduct := 0
	for i := 0; i < len(prob8Runes)-winSize; i++ {
		winProduct := 1
		for j := i; j < i+winSize; j++ {
			// Quick and dirty way of converting rune to corresponding integer
			winProduct *= int(prob8Runes[j] - '0')
			if winProduct > maxProduct {
				maxProduct = winProduct
			}
		}
	}
	return maxProduct
}

// Find a Pythagorean triplet where sum equals n.
// Proj Euler prob 9.
func pythagoreanTriplet(n int) (int, int, int, error) {
	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c, nil
			}
		}
	}
	return 0, 0, 0, errors.New("no satisfactory Pythagorean triplet found")
}

// Sum of all primes below n.
func sumPrimesBelowBound(n int) int {
	var sum int
	for val, isNotPrime := range *sieve(n) {
		if val >= 2 && !isNotPrime {
			sum += val
		}
	}
	return sum
}

var prob11GridRows = []string{
	"08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08",
	"49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00",
	"81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65",
	"52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91",
	"22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80",
	"24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50",
	"32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70",
	"67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21",
	"24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72",
	"21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95",
	"78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92",
	"16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57",
	"86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58",
	"19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40",
	"04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66",
	"88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69",
	"04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36",
	"20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16",
	"20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54",
	"01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48"}

// Find greatest product of winSize adjacent numbers in the same direction (up,
// down, left, right, or diagonally) in the prob11GridRows grid.
// Proj Euler prob 11.
func prob11(winSize int) int {
	ncols := len(strings.Split(prob11GridRows[0], " "))
	nrows := len(prob11GridRows)

	// Convert grid from slice of strings (one per row) to
	// 2D slice of integers (slice of slices)
	grid := make([][]int, ncols)
	for c := range grid {
		grid[c] = make([]int, nrows)
	}
	for y, row := range prob11GridRows {
		for x, val := range strings.Split(row, " ") {
			grid[x][y], _ = strconv.Atoi(val)
		}
	}

	var max int
	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if c < ncols-winSize {
				// Find product for runs of length winSize along rows
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r][c+i]
				}
				if prod > max {
					max = prod
				}
			}
			if r < nrows-winSize {
				// Find product for runs of length winSize down columns
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r+i][c]
				}
				if prod > max {
					max = prod
				}
			}
			if r < nrows-winSize && c < ncols-winSize {
				// Find product for diagonal runs of length winSize
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r+i][c+i]
				}
				if prod > max {
					max = prod
				}
			}
			if r > winSize-1 && c < ncols-winSize {
				// Find product for runs of length winSize along other diagonal
				prod := 1
				for i := 0; i < winSize; i++ {
					prod *= grid[r-i][c+i]
				}
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}

// Find the first triangle number with over n divisors.
// Proj Euler prob 12.
// Probably not a particularly efficient approach.
func minTriNumExcessOfNDivisors(n int) int {
	iter, tri := 1, 1

	for {
		nDivisors := 0
		floorSqrt := int(math.Floor(math.Sqrt(float64(tri))))
		for i := 1; i <= floorSqrt; i++ {
			if tri%i == 0 {
				if i == tri/i {
					nDivisors++
				} else {
					nDivisors += 2
				}
				if nDivisors > n {
					return tri
				}
			}
		}
		iter++
		tri += iter
	}
	// Shouldn't get here.
	return 0
}

var prob13Strs = []string{
	"37107287533902102798797998220837590246510135740250",
	"46376937677490009712648124896970078050417018260538",
	"74324986199524741059474233309513058123726617309629",
	"91942213363574161572522430563301811072406154908250",
	"23067588207539346171171980310421047513778063246676",
	"89261670696623633820136378418383684178734361726757",
	"28112879812849979408065481931592621691275889832738",
	"44274228917432520321923589422876796487670272189318",
	"47451445736001306439091167216856844588711603153276",
	"70386486105843025439939619828917593665686757934951",
	"62176457141856560629502157223196586755079324193331",
	"64906352462741904929101432445813822663347944758178",
	"92575867718337217661963751590579239728245598838407",
	"58203565325359399008402633568948830189458628227828",
	"80181199384826282014278194139940567587151170094390",
	"35398664372827112653829987240784473053190104293586",
	"86515506006295864861532075273371959191420517255829",
	"71693888707715466499115593487603532921714970056938",
	"54370070576826684624621495650076471787294438377604",
	"53282654108756828443191190634694037855217779295145",
	"36123272525000296071075082563815656710885258350721",
	"45876576172410976447339110607218265236877223636045",
	"17423706905851860660448207621209813287860733969412",
	"81142660418086830619328460811191061556940512689692",
	"51934325451728388641918047049293215058642563049483",
	"62467221648435076201727918039944693004732956340691",
	"15732444386908125794514089057706229429197107928209",
	"55037687525678773091862540744969844508330393682126",
	"18336384825330154686196124348767681297534375946515",
	"80386287592878490201521685554828717201219257766954",
	"78182833757993103614740356856449095527097864797581",
	"16726320100436897842553539920931837441497806860984",
	"48403098129077791799088218795327364475675590848030",
	"87086987551392711854517078544161852424320693150332",
	"59959406895756536782107074926966537676326235447210",
	"69793950679652694742597709739166693763042633987085",
	"41052684708299085211399427365734116182760315001271",
	"65378607361501080857009149939512557028198746004375",
	"35829035317434717326932123578154982629742552737307",
	"94953759765105305946966067683156574377167401875275",
	"88902802571733229619176668713819931811048770190271",
	"25267680276078003013678680992525463401061632866526",
	"36270218540497705585629946580636237993140746255962",
	"24074486908231174977792365466257246923322810917141",
	"91430288197103288597806669760892938638285025333403",
	"34413065578016127815921815005561868836468420090470",
	"23053081172816430487623791969842487255036638784583",
	"11487696932154902810424020138335124462181441773470",
	"63783299490636259666498587618221225225512486764533",
	"67720186971698544312419572409913959008952310058822",
	"95548255300263520781532296796249481641953868218774",
	"76085327132285723110424803456124867697064507995236",
	"37774242535411291684276865538926205024910326572967",
	"23701913275725675285653248258265463092207058596522",
	"29798860272258331913126375147341994889534765745501",
	"18495701454879288984856827726077713721403798879715",
	"38298203783031473527721580348144513491373226651381",
	"34829543829199918180278916522431027392251122869539",
	"40957953066405232632538044100059654939159879593635",
	"29746152185502371307642255121183693803580388584903",
	"41698116222072977186158236678424689157993532961922",
	"62467957194401269043877107275048102390895523597457",
	"23189706772547915061505504953922979530901129967519",
	"86188088225875314529584099251203829009407770775672",
	"11306739708304724483816533873502340845647058077308",
	"82959174767140363198008187129011875491310547126581",
	"97623331044818386269515456334926366572897563400500",
	"42846280183517070527831839425882145521227251250327",
	"55121603546981200581762165212827652751691296897789",
	"32238195734329339946437501907836945765883352399886",
	"75506164965184775180738168837861091527357929701337",
	"62177842752192623401942399639168044983993173312731",
	"32924185707147349566916674687634660915035914677504",
	"99518671430235219628894890102423325116913619626622",
	"73267460800591547471830798392868535206946944540724",
	"76841822524674417161514036427982273348055556214818",
	"97142617910342598647204516893989422179826088076852",
	"87783646182799346313767754307809363333018982642090",
	"10848802521674670883215120185883543223812876952786",
	"71329612474782464538636993009049310363619763878039",
	"62184073572399794223406235393808339651327408011116",
	"66627891981488087797941876876144230030984490851411",
	"60661826293682836764744779239180335110989069790714",
	"85786944089552990653640447425576083659976645795096",
	"66024396409905389607120198219976047599490197230297",
	"64913982680032973156037120041377903785566085089252",
	"16730939319872750275468906903707539413042652315011",
	"94809377245048795150954100921645863754710598436791",
	"78639167021187492431995700641917969777599028300699",
	"15368713711936614952811305876380278410754449733078",
	"40789923115535562561142322423255033685442488917353",
	"44889911501440648020369068063960672322193204149535",
	"41503128880339536053299340368006977710650566631954",
	"81234880673210146739058568557934581403627822703280",
	"82616570773948327592232845941706525094512325230608",
	"22918802058777319719839450180888072429661980811197",
	"77158542502016545090413245809786882778948721859617",
	"72107838435069186155435662884062257473692284509516",
	"20849603980134001723930671666823555245252804609722",
	"53503534226472524250874054075591789781264330331690",
}

// First 10 digits of sum of 100x 50-decimal-digit integers.
// Implemented using Go's native arbitrary precision support (big.Int).
// Proj Euler prob 13.
func prob13BigInt() (uint64, error) {
	accum := new(big.Int)
	for _, str := range prob13Strs {
		x := new(big.Int)
		x.SetString(str, 10)
		accum.Add(accum, x)
	}

	// Extract the first 10 decimal digits and convert to finite precision.
	retStr := string([]rune(fmt.Sprintf("%v", accum))[:10])
	return strconv.ParseUint(retStr, 10, 64)
}

func prob13CustomArbPrec() (uint64, error) {
	var sumLimbs uint64
	var carryOver uint64

	// Implement arbitrary precision support by splitting large integers into
	// runs of contiguous decimal digits ('limbs') of length 10.
	limbNDigits := 10
	// Number of decimal digits in all large integers in our input array.
	numNDigits := len([]rune(prob13Strs[0]))
	// After adding >=2 limbs divide the result by this to find what needs to be
	// carried over if moving on to process higher-order limbs.
	denominator := uint64(math.Pow(float64(10), float64(limbNDigits)))

	// For each limb range (from low-order to high-order):
	for i := numNDigits; i > 0; i -= limbNDigits {
		// Initally zero
		sumLimbs = carryOver

		// For each big integer in our input array:
		for _, s := range prob13Strs {
			// Extract the limb of interest as a uint64
			limb, err := strconv.ParseUint(string([]rune(s)[i-limbNDigits:i]), 10, 64)
			if err != nil {
				panic(err)
			}
			sumLimbs += limb
		}
		carryOver = sumLimbs / denominator
	}

	// Extract the first 10 decimal digits and convert to finite precision.
	retStr := string([]rune(fmt.Sprintf("%v", sumLimbs))[:10])
	return strconv.ParseUint(retStr, 10, 64)
}

// Length of Collatz seqence starting with n.
func collatzSeqLen(n uint) uint {
	var steps uint = 1
	for {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
		if n == 1 {
			break
		}
	}
	return steps
}

// Length of longest Collatz sequence for starting numbers 1 <= x < n.
// Proj Euler prob 14.
func longestCollatzSeqLen(n uint) uint {
	var start uint
	var maxChainLength uint
	var bestStart uint

	for start = 1; start < n; start++ {
		steps := collatzSeqLen(start)
		if steps > maxChainLength {
			maxChainLength = steps
			bestStart = start
		}
	}
	return bestStart
}

// Binomial Coefficient (n k).
//
// Here we find (n k) iteratively rather than using the standard approach of
//
//	n! / (k! * (n-k)!)
//
// so as to avoid overflows.
func binomCoeff(n uint64, k uint64) uint64 {
	var binomCoeff float64 = 1
	var i uint64

	for i < k {
		binomCoeff *= float64(n-i) / float64(k-i)
		i += 1
	}
	return uint64(binomCoeff)
}

// Number of different ways of moving from top left of square grid of width
// sideLen to bottom right.
//
// The problem is essential the binomial coefficient (n k),
// where n == 2 x sideLen and k == sideLen.
//
// Proj Euler prob 15.
func prob15(sideLen uint64) uint64 {
	return binomCoeff(sideLen*2, sideLen)
}

// Sum of digits of decimal n.
// Proj Euler prob 16.
func sumOfDigits(n uint) uint {
	var sum uint
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Sum of digits in the decimal representation of 2^n.
//
// NB need to represent decimal as arbitrary precision integers using an array
// so as to cope with larger values of n with less risk of overflows.
//
// Proj Euler prob 16.
func decDigitSum2PowerN(n uint) uint {
	// Upper bound on number of decimal digits in 2^n
	nDecDig := uint(math.Ceil(float64(n) * math.Log10(2)))

	// Array to store the digits of the decimal number
	arr := make([]uint, nDecDig)

	// Initally all zeros apart from the first element i.e. represents 2^0
	arr[0] = 1

	var i, j, tmp, carry, sum uint
	// For each power of 2:
	for i = 0; i < n; i++ {
		for j = 0; j < nDecDig; j++ {
			if carry == 0 && arr[j] == 0 {
				continue
			}
			tmp = (arr[j] * 2) + carry
			arr[j] = tmp % 10
			carry = tmp / 10
		}
	}
	for _, x := range arr {
		sum += x
	}
	return sum
}

var numWordMap = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// Express a decimal integer 1 <= n <= 1000 in English words.
func decimalToWords(n int) *[]string {
	if n > 1000 || n < 1 {
		fmt.Errorf("out of range: %v", n)
	}

	words := []string{}

	if n == 1000 {
		words = append(words, "one", "thousand")
		return &words
	}

	if n > 99 {
		digit := n / 100
		words = append(words, numWordMap[digit], "hundred")
		if n%100 == 0 {
			return &words
		}
		words = append(words, "and")
		n -= (100 * digit)
	}

	if n > 20 {
		words = append(words, numWordMap[(n/10)*10])
		if n%10 == 0 {
			return &words
		}
		words = append(words, numWordMap[n%10])
		return &words

	}
	words = append(words, numWordMap[n])
	return &words
}

// Sum of characters in all English words for the numbers 1..1000 inclusive.
// Excludes hyphens and spaces but not 'and'.
// Proj Euelr prob 17.
func prob17(lb int, ub int) int {
	var sum int

	for i := lb; i <= ub; i++ {
		for _, w := range *decimalToWords(i) {
			sum += utf8.RuneCountInString(w)
		}
	}
	return sum
}

// How many Sundays fell on the first of the month during the twentieth century
// (1 Jan 1901 to 31 Dec 2000)?
func prob19() int {
	var sum int = 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == time.Sunday {
				sum++
			}
		}
	}
	return sum
}

// factBigInt finds the factorial of an integer as a big.Int
//
// Helper for prob 20
func factBigInt(n int) big.Int {
	ret := big.NewInt(1)
	for i := 1; i <= n; i++ {
		ret.Mul(ret, big.NewInt(int64(i)))
	}
	return *ret
}

// sumDigitsBigInt finds the sum of all digits in a big.Int
//
// Helper for prob 20
func sumDigitsBigInt(n big.Int) int {
	sumDigits := 0
	for _, c := range n.String() {
		sumDigits += int(c - '0')
	}
	return sumDigits
}

func prob20() int {
	return sumDigitsBigInt(factBigInt(100))
}

// sumNameScores sorts a string of names, finds the sum of the alphabetical
// positions of each character in each name, finds the product of that sum and
// the name's position in the list, then finally returns the sum of those
// products.
//
// namesStr is a string of comma-delimited names where each is encapsulated in
// double quotes.
//
// Solution to prob 22
func sumNameScores(namesStr string) int {
	namesStrLower := strings.ToLower(namesStr)
	names := strings.Split(namesStrLower, ",")
	sort.Strings(names)

	finalSum := 0
	for i, n := range names {
		orderScore := i + 1

		alphaScore := 0
		for _, c := range strings.Trim(n, "\"") {
			// Add the position of each character in the latin alphabet
			// (a: 1, b: 2, etc)
			alphaScore += int(byte(c)-byte('a')) + 1
		}
		finalSum += orderScore * alphaScore
	}
	return finalSum
}

func properDivisors(n int) []int {
	// Proper divisors of n.
    // 
	// Not guaranteed to be in sorted order.
	// NB the proper divisors of a number include 1 but not the number itself.
	divs := []int{1}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n % i == 0 {
			divs = append(divs, i)
			j := n / i
			if i != j {
				divs = append(divs, j)
			}
		}
	}
	return divs
}

func isAbundant(n int) bool {
	// An abundant numbe is one where the sum of the proper divisors 
	// is greater than the number itself.
	var spds int
	for _, i := range(properDivisors(n)) {
		spds += i
	}
	return spds > n
}

func prob23(upperBound int) int {
	// Find the sum of all numbers that aren't the sum of two abundant numbers

	// Build a slice of all abundant numbers below an upper limit
	var abundants []int
	for i := 1; i <= upperBound; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}
	// Build a set all numbers that are the sum of two abundant numbers
	// using abundant numbers from that slice
	sumTwoAbundants := make(map[int]bool)
	for _, i := range abundants {
		for _, j := range abundants {
			sumTwoAbundants[i+j] = true
		}
	}
	// Find the sum of all numbers that aren't the sum of two abundant numbers
	sum := 0
	for i := 1; i <= upperBound; i++ {
		if _, isSumTwoAbundants := sumTwoAbundants[i]; !isSumTwoAbundants {
			sum += i
		}
	}
	return sum
}

func main() {
	//type BinaryNode struct {
	//	left  *BinaryNode
	//	right *BinaryNode
	//	data  int
	//	id    int
	//}
	//type BinaryTree struct {
	//	root *BinaryNode
	//}
	// See https://www.golangprograms.com/golang-program-to-implement-binary-tree.html for more implementation info.
	//
	// Then perform Dijkstra's algorithm on this binary tree
}

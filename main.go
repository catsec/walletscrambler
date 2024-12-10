package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func wordExists(word string, wordlist []string) bool {
	for _, w := range wordlist {
		if w == word {
			return true
		}
	}
	return false
}

func isWeakPassword(password string) bool {
	if len(password) < 8 {
		return true
	}

	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasNumber := strings.ContainsAny(password, "0123456789")
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:',.<>?/")

	return !(hasLower && hasUpper && hasNumber && hasSpecial)
}

func hashRepeatedly(data []byte, iterations int) []byte {
	hash := data
	for i := 0; i < iterations; i++ {
		digest := sha3.Sum256(hash)
		hash = digest[:]
	}
	return hash
}
func bytesToBitString(data []byte) string {
	var bitString strings.Builder
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			if (b & (1 << i)) != 0 {
				bitString.WriteByte('1')
			} else {
				bitString.WriteByte('0')
			}
		}
	}
	return bitString.String()
}
func splitString(input string, length int) []string {
	if length <= 0 {
		return []string{}
	}

	var result []string
	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}

func xorBitStrings(bits1, bits2 string) string {
	if len(bits1) < len(bits2) {
		bits1 = strings.Repeat("0", len(bits2)-len(bits1)) + bits1
	} else if len(bits2) < len(bits1) {
		bits2 = strings.Repeat("0", len(bits1)-len(bits2)) + bits2
	}

	var result strings.Builder
	for i := 0; i < len(bits1); i++ {
		if bits1[i] == bits2[i] {
			result.WriteByte('0')
		} else {
			result.WriteByte('1')
		}
	}

	return result.String()
}

func bitsToInt(bits string) int {
	result := new(big.Int)
	result.SetString(bits, 2)
	return int(result.Int64())
}

func intToBits(value int, bitLength int) string {
	bitString := strconv.FormatInt(int64(value), 2)
	if len(bitString) < bitLength {
		padding := bitLength - len(bitString)
		bitString = strings.Repeat("0", padding) + bitString
	}

	return bitString
}
func printBeautifully(title string, words []string) {
	fmt.Printf("\n%s\n%s\n", title, strings.Repeat("=", len(title)))
	numberPadding := len(fmt.Sprintf("%d", len(words)))
	for i, word := range words {
		fmt.Printf("%*d. %s\n", numberPadding, i+1, word)
	}
}
func main() {
	var words = []string{
		"academic", "acid", "acne", "acquire", "acrobat", "activity", "actress", "adapt", "adequate", "adjust",
		"admit", "adorn", "adult", "advance", "advocate", "afraid", "again", "agency", "agree", "aide",
		"aircraft", "airline", "airport", "ajar", "alarm", "album", "alcohol", "alien", "alive", "alpha",
		"already", "alto", "aluminum", "always", "amazing", "ambition", "amount", "amuse", "analysis", "anatomy",
		"ancestor", "ancient", "angel", "angry", "animal", "answer", "antenna", "anxiety", "apart", "aquatic",
		"arcade", "arena", "argue", "armed", "artist", "artwork", "aspect", "auction", "august", "aunt",
		"average", "aviation", "avoid", "award", "away", "axis", "axle", "beam", "beard", "beaver",
		"become", "bedroom", "behavior", "being", "believe", "belong", "benefit", "best", "beyond", "bike",
		"biology", "birthday", "bishop", "black", "blanket", "blessing", "blimp", "blind", "blue", "body",
		"bolt", "boring", "born", "both", "boundary", "bracelet", "branch", "brave", "breathe", "briefing",
		"broken", "brother", "browser", "bucket", "budget", "building", "bulb", "bulge", "bumpy", "bundle",
		"burden", "burning", "busy", "buyer", "cage", "calcium", "camera", "campus", "canyon", "capacity",
		"capital", "capture", "carbon", "cards", "careful", "cargo", "carpet", "carve", "category", "cause",
		"ceiling", "center", "ceramic", "champion", "change", "charity", "check", "chemical", "chest", "chew",
		"chubby", "cinema", "civil", "class", "clay", "cleanup", "client", "climate", "clinic", "clock",
		"clogs", "closet", "clothes", "club", "cluster", "coal", "coastal", "coding", "column", "company",
		"corner", "costume", "counter", "course", "cover", "cowboy", "cradle", "craft", "crazy", "credit",
		"cricket", "criminal", "crisis", "critical", "crowd", "crucial", "crunch", "crush", "crystal", "cubic",
		"cultural", "curious", "curly", "custody", "cylinder", "daisy", "damage", "dance", "darkness", "database",
		"daughter", "deadline", "deal", "debris", "debut", "decent", "decision", "declare", "decorate", "decrease",
		"deliver", "demand", "density", "deny", "depart", "depend", "depict", "deploy", "describe", "desert",
		"desire", "desktop", "destroy", "detailed", "detect", "device", "devote", "diagnose", "dictate", "diet",
		"dilemma", "diminish", "dining", "diploma", "disaster", "discuss", "disease", "dish", "dismiss", "display",
		"distance", "dive", "divorce", "document", "domain", "domestic", "dominant", "dough", "downtown", "dragon",
		"dramatic", "dream", "dress", "drift", "drink", "drove", "drug", "dryer", "duckling", "duke",
		"duration", "dwarf", "dynamic", "early", "earth", "easel", "easy", "echo", "eclipse", "ecology",
		"edge", "editor", "educate", "either", "elbow", "elder", "election", "elegant", "element", "elephant",
		"elevator", "elite", "else", "email", "emerald", "emission", "emperor", "emphasis", "employer", "empty",
		"ending", "endless", "endorse", "enemy", "energy", "enforce", "engage", "enjoy", "enlarge", "entrance",
		"envelope", "envy", "epidemic", "episode", "equation", "equip", "eraser", "erode", "escape", "estate",
		"estimate", "evaluate", "evening", "evidence", "evil", "evoke", "exact", "example", "exceed", "exchange",
		"exclude", "excuse", "execute", "exercise", "exhaust", "exotic", "expand", "expect", "explain", "express",
		"extend", "extra", "eyebrow", "facility", "fact", "failure", "faint", "fake", "false", "family",
		"famous", "fancy", "fangs", "fantasy", "fatal", "fatigue", "favorite", "fawn", "fiber", "fiction",
		"filter", "finance", "findings", "finger", "firefly", "firm", "fiscal", "fishing", "fitness", "flame",
		"flash", "flavor", "flea", "flexible", "flip", "float", "floral", "fluff", "focus", "forbid",
		"force", "forecast", "forget", "formal", "fortune", "forward", "founder", "fraction", "fragment", "frequent",
		"freshman", "friar", "fridge", "friendly", "frost", "froth", "frozen", "fumes", "funding", "furl",
		"fused", "galaxy", "game", "garbage", "garden", "garlic", "gasoline", "gather", "general", "genius",
		"genre", "genuine", "geology", "gesture", "glad", "glance", "glasses", "glen", "glimpse", "goat",
		"golden", "graduate", "grant", "grasp", "gravity", "gray", "greatest", "grief", "grill", "grin",
		"grocery", "gross", "group", "grownup", "grumpy", "guard", "guest", "guilt", "guitar", "gums",
		"hairy", "hamster", "hand", "hanger", "harvest", "have", "havoc", "hawk", "hazard", "headset",
		"health", "hearing", "heat", "helpful", "herald", "herd", "hesitate", "hobo", "holiday", "holy",
		"home", "hormone", "hospital", "hour", "huge", "human", "humidity", "hunting", "husband", "hush",
		"husky", "hybrid", "idea", "identify", "idle", "image", "impact", "imply", "improve", "impulse",
		"include", "income", "increase", "index", "indicate", "industry", "infant", "inform", "inherit", "injury",
		"inmate", "insect", "inside", "install", "intend", "intimate", "invasion", "involve", "iris", "island",
		"isolate", "item", "ivory", "jacket", "jerky", "jewelry", "join", "judicial", "juice", "jump",
		"junction", "junior", "junk", "jury", "justice", "kernel", "keyboard", "kidney", "kind", "kitchen",
		"knife", "knit", "laden", "ladle", "ladybug", "lair", "lamp", "language", "large", "laser",
		"laundry", "lawsuit", "leader", "leaf", "learn", "leaves", "lecture", "legal", "legend", "legs",
		"lend", "length", "level", "liberty", "library", "license", "lift", "likely", "lilac", "lily",
		"lips", "liquid", "listen", "literary", "living", "lizard", "loan", "lobe", "location", "losing",
		"loud", "loyalty", "luck", "lunar", "lunch", "lungs", "luxury", "lying", "lyrics", "machine",
		"magazine", "maiden", "mailman", "main", "makeup", "making", "mama", "manager", "mandate", "mansion",
		"manual", "marathon", "march", "market", "marvel", "mason", "material", "math", "maximum", "mayor",
		"meaning", "medal", "medical", "member", "memory", "mental", "merchant", "merit", "method", "metric",
		"midst", "mild", "military", "mineral", "minister", "miracle", "mixed", "mixture", "mobile", "modern",
		"modify", "moisture", "moment", "morning", "mortgage", "mother", "mountain", "mouse", "move", "much",
		"mule", "multiple", "muscle", "museum", "music", "mustang", "nail", "national", "necklace", "negative",
		"nervous", "network", "news", "nuclear", "numb", "numerous", "nylon", "oasis", "obesity", "object",
		"observe", "obtain", "ocean", "often", "olympic", "omit", "oral", "orange", "orbit", "order",
		"ordinary", "organize", "ounce", "oven", "overall", "owner", "paces", "pacific", "package", "paid",
		"painting", "pajamas", "pancake", "pants", "papa", "paper", "parcel", "parking", "party", "patent",
		"patrol", "payment", "payroll", "peaceful", "peanut", "peasant", "pecan", "penalty", "pencil", "percent",
		"perfect", "permit", "petition", "phantom", "pharmacy", "photo", "phrase", "physics", "pickup", "picture",
		"piece", "pile", "pink", "pipeline", "pistol", "pitch", "plains", "plan", "plastic", "platform",
		"playoff", "pleasure", "plot", "plunge", "practice", "prayer", "preach", "predator", "pregnant", "premium",
		"prepare", "presence", "prevent", "priest", "primary", "priority", "prisoner", "privacy", "prize", "problem",
		"process", "profile", "program", "promise", "prospect", "provide", "prune", "public", "pulse", "pumps",
		"punish", "puny", "pupal", "purchase", "purple", "python", "quantity", "quarter", "quick", "quiet",
		"race", "racism", "radar", "railroad", "rainbow", "raisin", "random", "ranked", "rapids", "raspy",
		"reaction", "realize", "rebound", "rebuild", "recall", "receiver", "recover", "regret", "regular", "reject",
		"relate", "remember", "remind", "remove", "render", "repair", "repeat", "replace", "require", "rescue",
		"research", "resident", "response", "result", "retailer", "retreat", "reunion", "revenue", "review", "reward",
		"rhyme", "rhythm", "rich", "rival", "river", "robin", "rocky", "romantic", "romp", "roster",
		"round", "royal", "ruin", "ruler", "rumor", "sack", "safari", "salary", "salon", "salt",
		"satisfy", "satoshi", "saver", "says", "scandal", "scared", "scatter", "scene", "scholar", "science",
		"scout", "scramble", "screw", "script", "scroll", "seafood", "season", "secret", "security", "segment",
		"senior", "shadow", "shaft", "shame", "shaped", "sharp", "shelter", "sheriff", "short", "should",
		"shrimp", "sidewalk", "silent", "silver", "similar", "simple", "single", "sister", "skin", "skunk",
		"slap", "slavery", "sled", "slice", "slim", "slow", "slush", "smart", "smear", "smell",
		"smirk", "smith", "smoking", "smug", "snake", "snapshot", "sniff", "society", "software", "soldier",
		"solution", "soul", "source", "space", "spark", "speak", "species", "spelling", "spend", "spew",
		"spider", "spill", "spine", "spirit", "spit", "spray", "sprinkle", "square", "squeeze", "stadium",
		"staff", "standard", "starting", "station", "stay", "steady", "step", "stick", "stilt", "story",
		"strategy", "strike", "style", "subject", "submit", "sugar", "suitable", "sunlight", "superior", "surface",
		"surprise", "survive", "sweater", "swimming", "swing", "switch", "symbolic", "sympathy", "syndrome", "system",
		"tackle", "tactics", "tadpole", "talent", "task", "taste", "taught", "taxi", "teacher", "teammate",
		"teaspoon", "temple", "tenant", "tendency", "tension", "terminal", "testify", "texture", "thank", "that",
		"theater", "theory", "therapy", "thorn", "threaten", "thumb", "thunder", "ticket", "tidy", "timber",
		"timely", "ting", "tofu", "together", "tolerate", "total", "toxic", "tracks", "traffic", "training",
		"transfer", "trash", "traveler", "treat", "trend", "trial", "tricycle", "trip", "triumph", "trouble",
		"true", "trust", "twice", "twin", "type", "typical", "ugly", "ultimate", "umbrella", "uncover",
		"undergo", "unfair", "unfold", "unhappy", "union", "universe", "unkind", "unknown", "unusual", "unwrap",
		"upgrade", "upstairs", "username", "usher", "usual", "valid", "valuable", "vampire", "vanish", "various",
		"vegan", "velvet", "venture", "verdict", "verify", "very", "veteran", "vexed", "victim", "video",
		"view", "vintage", "violence", "viral", "visitor", "visual", "vitamins", "vocal", "voice", "volume",
		"voter", "voting", "walnut", "warmth", "warn", "watch", "wavy", "wealthy", "weapon", "webcam",
		"welcome", "welfare", "western", "width", "wildlife", "window", "wine", "wireless", "wisdom", "withdraw",
		"wits", "wolf", "woman", "work", "worthy", "wrap", "wrist", "writing", "wrote", "year",
		"yelp", "yield", "yoga", "zero",
	}

	wordCount := len(words)

	if !isPowerOfTwo(wordCount) {
		fmt.Printf("Error: The wordlist must contain a power of 2 elements, but it contains %d elements.", wordCount)
		return
	}
	fmt.Print("Welcome to the wallet word scrambler\n\n")
	fmt.Print("It will use a password and salt to create a new wallet words backup\n")
	fmt.Print("from your existing wallet words\n\n")
	fmt.Print("The word list is the SLIP39 English wordlist containing 1024 words\n\n")
	fmt.Print("Warning:\n")
	fmt.Print("This program is meant to run on a formated and air gapped machine\n")
	fmt.Print("It is not safe to run it on a machine connected to any kind of network\n")
	fmt.Print("Even tough we save nothing - wipe your machine after use\n\n")
	reader := bufio.NewReader(os.Stdin)
	var password1, password2 string
	for {
		fmt.Print("Enter your password: ")
		password1, _ = reader.ReadString('\n')
		password1 = strings.TrimSpace(password1)

		fmt.Print("Confirm the password: ")
		password2, _ = reader.ReadString('\n')
		password2 = strings.TrimSpace(password2)

		if password1 != password2 {
			fmt.Println("Error: Passwords do not match. Try again.")
			continue
		}

		if isWeakPassword(password1) {
			fmt.Println("Warning: Your password is weak. It should be at least 8 characters long and include a mix of uppercase, lowercase, numbers, and special characters.")
			fmt.Print("Do you want to continue with this password? (yes/no): ")
			confirmation, _ := reader.ReadString('\n')
			confirmation = strings.TrimSpace(strings.ToLower(confirmation))
			if confirmation == "yes" {
				fmt.Println("Password accepted.")
				break
			} else {
				fmt.Println("Please enter a stronger password.")
				continue
			}
		} else {
			fmt.Println("Password accepted.")
			break
		}
	}
	fmt.Print("\n *** NEVER forget your password - there is no way to recover it ***\n\n")
	var saltOption string
	fmt.Println("Do you have salt words or would you like to generate them randomly?\n('(M)anual' or '(R)andom'):")
	for {
		fmt.Print("Your choice: ")
		saltOption, _ = reader.ReadString('\n')
		saltOption = strings.TrimSpace(strings.ToLower(saltOption))
		if saltOption == "m" || saltOption == "r" {
			break
		}
		fmt.Println("Invalid choice. Please enter 'M' or 'R'.")
	}

	var saltCount int

	for {
		fmt.Print("\nEnter the number of salt words (0-16, at least 4 recommended): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil || number < 0 || number > 16 {
			fmt.Println("\nInvalid input. Please enter a number between 0 and 16.")
			continue
		}
		saltCount = number
		break
	}

	var saltWords []string

	if saltOption == "m" {
		for i := 0; i < saltCount; i++ {
			for {
				fmt.Printf("Enter salt word %d: ", i+1)
				word, _ := reader.ReadString('\n')
				word = strings.TrimSpace(word)
				if !wordExists(word, words) {
					fmt.Println("Invalid word. The word must exist in the wordlist.")
				} else {
					saltWords = append(saltWords, word)
					break
				}
			}
		}
	} else if saltOption == "r" {
		for i := 0; i < saltCount; i++ {
			index, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
			if err != nil {
				fmt.Printf("Error generating random index: %v", err)
				return
			}
			randomWord := words[index.Int64()]
			saltWords = append(saltWords, randomWord)
		}
	}

	fmt.Print("\nCalculating key from your salt and password.\n")
	fmt.Print("For safty reasons, this is supposed to take a while...\n\n")

	salt := strings.Join(saltWords, "")
	argon2Seed := hashRepeatedly([]byte(salt), 4847868)

	memory := uint32(1024 * 1024)
	time := uint32(64)
	threads := uint8(4)
	keyLen := uint32(64)

	argon2Hash := argon2.IDKey([]byte(password1), argon2Seed, time, memory, threads, keyLen)
	varkeybits := bytesToBitString(argon2Hash)
	wordBitSize := int(math.Log2(float64(wordCount)))
	keybitswords := splitString(varkeybits, wordBitSize)

	var walletWordCount int
	for {
		fmt.Print("\nEnter the number of words in your wallet (4-33): ")
		_, err := fmt.Scan(&walletWordCount)
		if err == nil && walletWordCount >= 4 && walletWordCount <= 33 {
			break
		}
		fmt.Println("Invalid input. Please enter a number between 12 and 33.")
	}

	newWords := make([]string, walletWordCount)
	for i := 0; i < walletWordCount; i++ {
		var word string
		for {
			fmt.Printf("Enter word %d: ", i+1)
			word, _ = reader.ReadString('\n')
			word = strings.TrimSpace(word)
			if wordExists(word, words) {
				break
			}
			fmt.Println("Invalid word. Please enter a valid word from the wordlist.")
		}

		wordIndex := -1
		for j, w := range words {
			if w == word {
				wordIndex = j
				break
			}
		}

		wordBits := intToBits(wordIndex, wordBitSize)
		xorResult := xorBitStrings(keybitswords[i], wordBits)
		newWordIndex := bitsToInt(xorResult)
		newWords[i] = words[newWordIndex]
	}

	fmt.Print("\n*** Here are your new wallet words ***\n")

	printBeautifully("Salt:", saltWords)
	printBeautifully("Wallet Words:", newWords)

	fmt.Print("\n\nWrite both salt and words down and store them in a safe place.\n\n")
	fmt.Print("*** NEVER forget your password - there is no way to recover it ***\n\n")
}

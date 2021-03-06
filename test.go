package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Test tests functions of go-chess
func Test() {
	fmt.Println("\nTesting Boards: ")

	fmt.Println("Board structures:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq120ToSq64[index])
	}
	fmt.Println()
	fmt.Println()
	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq64ToSq120[index])
	}

	fmt.Println("\nTesting Bitboards: ")

	fmt.Println("\nCounting and popping of pawns:")
	var playBitBoard Bitboard
	playBitBoard = 0
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D2)))
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D3)))
	playBitBoard |= (Bitboard(1) << Bitboard(SQ64(D4)))
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i := (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())
	i = (&playBitBoard).Pop()
	fmt.Println("Index: ", i)
	playBitBoard.Print()
	fmt.Println("\nCount: ", playBitBoard.Count())

	fmt.Println("Setting and clearing of bits on the chess board:")
	playBitBoard = 0
	(&playBitBoard).Set(61)
	playBitBoard.Print()
	(&playBitBoard).Clear(61)
	playBitBoard.Print()

	fmt.Println("\nTesting Parsing of FEN notation: ")

	var board Board

	// Position of pieces
	fmt.Println("\nInitial board structure: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err := ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nNumbers represent number of blank squares: 4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1")
	err = ParseFEN("4kbnr/5p2/8/8/8/8/PPPP4/RNB2BNR w KQkq a4 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nNumber should be 1-8: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("\n\nInitial board structure with a missing /: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = ParseFEN("rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	// Side
	fmt.Println("\n\nInitial board structure with missing side to play: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\n\nError when given invalid side: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR s KQkq a4 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Castling
	fmt.Println("\n\nInitial board structure with all castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	fmt.Println("\n\nInitial board structure with few castling permissions: rnbqkbnrpppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	fmt.Println("\n\nMissing castling condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Castling permissions: ", board.CastlePerm)

	// EnPas
	fmt.Println("\n\nChecking enPas when nil: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\n\nChecking enPas: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a1 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\n\nMissing enPas condition: rnbqkbnr/pppppppp/0/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq ", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("EnPas Condition: ", board.EnPas)

	fmt.Println("\nAfter first move: rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	err = ParseFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nAfter second move: rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	err = ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nAfter third move: rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	err = ParseFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	fmt.Println("\nFiles Board:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 && index != 0 {
			fmt.Println()
		}
		fmt.Printf("%4d", FilesBrd[index])
	}

	fmt.Println("\nRanks Board:")
	for index := 0; index < BrdSqNum; index++ {
		if index%10 == 0 && index != 0 {
			fmt.Println()
		}
		fmt.Printf("%4d", RanksBrd[index])
	}

	fmt.Println("\nPrinting bitboards of pawns:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.Print()
	fmt.Println("\nWhite pawn: ")
	pawnStructure := Bitboard(board.Pawns[White])
	pawnStructure.Print()
	fmt.Println("\nBlack pawn: ")
	pawnStructure = Bitboard(board.Pawns[Black])
	pawnStructure.Print()
	fmt.Println("\nBoth pawns: ")
	pawnStructure = Bitboard(board.Pawns[Both])
	pawnStructure.Print()

	fmt.Println("\nChecking if board structure is correct:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}

	fmt.Println("\nChecking if board structure is incorrect after forcefully changing piece count:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.PceNum[Wp]--
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}

	fmt.Println("\nChecking if board structure is incorrect after forcefully changing posKey:")
	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	board.PosKey ^= SideKey
	err = (&board).Check()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Checking successful")
	}

	// fmt.Println("\nChecking squares attacked when FEN is 8/3q4/8/8/4Q3/8/8/8 w - - 0 2:")
	// ParseFEN("8/3q4/8/8/4Q3/8/8/8 w - - 0 2", &board)
	// board.Print()
	// fmt.Println("\nWhite attacking: ")
	// printSqAttacked(White, board)
	// fmt.Println("\nBlack attacking: ")
	// printSqAttacked(Black, board)
	//
	// fmt.Println("\nChecking squares attacked when FEN is 8/3q1p2/8/5P2/4Q3/8/8/8 w - - 0 2:")
	// ParseFEN("8/3q1p2/8/5P2/4Q3/8/8/8 w - - 0 2", &board)
	// board.Print()
	// fmt.Println("\nWhite attacking: ")
	// printSqAttacked(White, board)
	// fmt.Println("\nBlack attacking: ")
	// printSqAttacked(Black, board)

	fmt.Println("Check from and to squares of move: 0110001111011010")
	num, _ := strconv.ParseInt("0110001111011010", 2, 32)

	result := FromSq(int(num))
	fmt.Println("From square: ", strconv.FormatInt(int64(result), 16))

	result = ToSq(int(num))
	fmt.Println("To square: ", strconv.FormatInt(int64(result), 16))

	fmt.Println("\nMaking a move with from = 6, to = 12, capture = Wr, promote = Br: ")
	move := 0
	from := 6
	to := 12
	capture := Wr
	promote := Br
	move = from | (to << 7) | (capture << 14) | (promote << 20)
	fmt.Printf("\nDecimal: %d, Hex: %s, Binary: %s\n", move, strconv.FormatInt(int64(move), 16), strconv.FormatInt(int64(move), 2))
	fmt.Printf("\nChecking functions with move as input: From: %d To: %d Captured: %d, Promoted: %d\n", FromSq(move), ToSq(move), Captured(move), Promoted(move))
	fmt.Println("Not Added flag for pawn start: ", (move & MFlagPS))
	move |= MFlagPS
	fmt.Println("Added flag for pawn start: ", (move & MFlagPS))

	fmt.Println("\nPrinting full move with promotion character as from = A2, to = H7, capture = Wr, promote = Wb:")
	move = A2 | (H7 << 7) | (Wr << 14) | (Wb << 20)
	fmt.Printf("Algebraic from: %s\n", PrintSq(A2))
	fmt.Printf("Algebraic to: %s\n", PrintSq(H7))
	fmt.Printf("Algebraic move: %s\n", PrintMove(move))

	fmt.Println("\nPrinting full move with no promotion character as from = A2, to = H7, capture = Wr")
	move = A2 | (H7 << 7) | (Wr << 14)
	fmt.Printf("Algebraic from: %s\n", PrintSq(A2))
	fmt.Printf("Algebraic to: %s\n", PrintSq(H7))
	fmt.Printf("Algebraic move: %s\n", PrintMove(move))

	fmt.Println("\nPrinting full move with invalid promotion character as from = A2, to = H7, capture = Wr, promote = Bk")
	move = A2 | (H7 << 7) | (Wr << 14) | (Bk << 20)
	fmt.Printf("Algebraic from: %s\n", PrintSq(A2))
	fmt.Printf("Algebraic to: %s\n", PrintSq(H7))
	fmt.Printf("Algebraic move: %s\n", PrintMove(move))

	fmt.Println("\nTesting total white pawn moves: rnbqkb1r/pp1p1pPp/8/2p1pR2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1")
	err = ParseFEN("rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	var movelist MoveList
	(&movelist).GenerateAllMoves(&board)
	(&movelist).Print()

	fmt.Println("\nTesting total black pawn moves: rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq e3 0 1")
	err = ParseFEN("rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq e3 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()

	(&movelist).GenerateAllMoves(&board)
	(&movelist).Print()

	fmt.Println("Testing moves of sliding and non-sliding pieces:")
	err = ParseFEN("rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of sliding and non-sliding pieces:")
	err = ParseFEN("rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq e3 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of black knights:")
	err = ParseFEN("5k2/1n6/4n3/6N1/8/3N4/8/5K2 b - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of white knights:")
	err = ParseFEN("5k2/1n6/4n3/6N1/8/3N4/8/5K2 w - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of white rook:")
	err = ParseFEN("6k1/8/5r2/8/1nR5/5N2/8/6K1 w - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of black rook:")
	err = ParseFEN("6k1/8/5r2/8/1nR5/5N2/8/6K1 b - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of white queen:")
	err = ParseFEN("6k1/8/4nq2/8/1nQ5/5N2/1N6/6K1 w - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of black queen:")
	err = ParseFEN("6k1/8/4nq2/8/1nQ5/5N2/1N6/6K1 b - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of black bishops:")
	err = ParseFEN("6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 b - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves of white bishops:")
	err = ParseFEN("6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 w - - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves in black side castling position:")
	err = ParseFEN("r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves in white side castling position:")
	err = ParseFEN("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves in white side castling position when intermediate squares are attacked:")
	err = ParseFEN("3rk2r/8/8/8/8/8/6p1/R3K2R w KQk - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing moves in black side castling position when a piece has been moved:")
	err = ParseFEN("3rk2r/8/8/8/8/8/6p1/R3K2R b KQk - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)

	fmt.Println("Testing move generation list:")
	err = ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)
	(&movelist).Print()

	fmt.Println("Testing make move and take move:")
	err = ParseFEN(StartFEN, &board)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	(&board).Print()
	(&movelist).GenerateAllMoves(&board)
	(&movelist).Print()
	for i := 0; i < movelist.count; i++ {
		moveMade, err := MakeMove(movelist.moves[i].move, &board)
		if !moveMade {
			fmt.Println("Error occurred: ", err)
			continue
		}
		fmt.Println("Move made: ", PrintMove(movelist.moves[i].move))
		board.Print()
		TakeMove(&board)
		fmt.Println("Move taken: ", PrintMove(movelist.moves[i].move))
		board.Print()
	}

	fmt.Println("Perft Testing")
	file, err := os.Open("perftsuite.md")
	if err != nil {
		log.Fatal("Error in opening file: ", err)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	for scanner.Scan() {
		testArr := strings.Split(scanner.Text(), ";")
		var board1 Board
		fmt.Println("Fen: ", testArr[0])
		err = ParseFEN(testArr[0], &board1)
		if err != nil {
			fmt.Println("Error in parsing fen: ", err)
		}
		for i := 1; i < len(testArr); i++ {
			testConditionArr := strings.Split(testArr[i], " ")
			depthArr := re.FindAllString(testConditionArr[0], -1)
			depth, _ := strconv.Atoi(depthArr[0])
			PerftTest(depth, &board1)
			fmt.Println("Expected result: ", testConditionArr[1])
		}
	}

	fmt.Println("Checking if program searches for mate in 3 position")
	mateIn3FEN := "2rr3k/pp3pp1/1nnqbN1p/3pN3/2pP4/2P3Q1/PPB4P/R4RK1 w - -"
	var board2 Board
	var info SearchInfo
	err = ParseFEN(mateIn3FEN, &board2)
	if err != nil {
		fmt.Println("Error in parsing fen: ", err)
	}
	board2.Print()
	(&board2).PvTable.Init()
	info.depth = 4
	SearchPosition(&board2, &info)
}

func printSqAttacked(side int, pos Board) {
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			sq := FR2SQ(file, rank)
			result, err1 := SqAttacked(sq, side, &pos)
			if err1 != nil {
				fmt.Println("Error: ", err1)
			}
			if result {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}

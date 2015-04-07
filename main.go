package main

import (
	"bufio"
	"github.com/biogo/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/seq/linear"
	"fmt"
	"os"
)

//Usage: RevComp inputFileName outputFileName
func main() {

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// open inputFile in read-only mode
	inputFile, inputError := os.Open(inputFileName)
	if inputError != nil {
		fmt.Println("An error occurred opening the input file.")
		return
	}
	defer inputFile.Close()

	// open outputFile in write-only mode
	outputFile, outputError := os.Create(outputFileName)
	if outputError != nil {
		fmt.Println("An error occurred with file creation.")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	sc := seqio.NewScanner(fasta.NewReader(inputFile, linear.NewSeq("", nil, alphabet.DNA)))
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)
		for i := range s.Seq {
			s.Seq[i] &^= ' '
		}
		// indicate that this is the reverse complemented sequence
		s.ID = s.ID + ".rev"
		// write header
		outputWriter.WriteString(">" + s.ID + "\n")
		// reverse complement the sequence
		s.RevComp()
		// write FASTA sequence
		outputWriter.WriteString(s.String() + "\n")
	}
	if sc.Error() != nil {
		panic(sc.Error())
	}

	// write the buffer completely to the file
	outputWriter.Flush()
}

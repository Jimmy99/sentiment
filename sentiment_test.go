package sentiment

import "testing"
import "fmt"

const aliceText = "There were doors all round the hall, but they were all locked, and when Alice had been all the way down one side and up the other, trying every door, she walked sadly down the middle, wondering how she was ever to get out again. Suddenly she came upon a little three-legged table, all made of solid glass; there was nothing on it but a tiny golden key, and Alice's first idea was that this might belong to one of the doors of the hall; but alas! either the locks were too large, or the key was too small, but at any rate it would not open any of them. However, on the second time round, she came upon a low curtain she had not noticed before, and behind it was a little door about fifteen inches high; she tried the little golden key in the lock, and to her great delight it fitted!"

func TestExclusion(t *testing.T) {
    h := NewHistogram("Jake")
    h.Absorb([]string{"Jake","Prat", "Prat"})
    fmt.Print(h.String())
        
    _, ok := h.Freq["Jake"] 
    if ok {
        t.Error("Jake should be exluded.")
    }

    freq, ok := h.Freq["Prat"]
    if !ok {
        t.Error("Prat should be included.")
    }

    if freq != 2 {
        t.Error("Prat should be counted twice.")
    }
}

func TestBigInput(t *testing.T) {
    h := NewHistogram()
    h.AbsorbText(aliceText, " ")
    fmt.Print(h.String())
}

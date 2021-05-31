package questions;


import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class GuessNumberHigherOrLowerTest {
    @Test
    public void TestMethod1() {
        GuessNumberHigherOrLower outer = new GuessNumberHigherOrLower();
        GuessNumberHigherOrLower.Solution solution = outer.new Solution(6);
        Assertions.assertEquals(6, solution.guessNumber(10));

    }
}
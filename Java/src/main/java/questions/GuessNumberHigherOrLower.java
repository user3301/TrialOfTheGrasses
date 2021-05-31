package questions;

public class GuessNumberHigherOrLower {

    public class GuessGame
    {
        private int pick;

        public GuessGame(int pick) {
            this.pick = pick;
        }

        public int guess(int num) {
            return num == pick ? 0 : (num > pick ? -1 : 1);
        }
    }

    public class Solution extends GuessGame{

        public Solution(int pick) {
            super(pick);
        }

        public int guessNumber(int n)  {
            int i = 1, j = n;
            while(i < j) {
                int mid = i + (j - i) / 2;
                if(this.guess(mid) == 0) {
                    return mid;
                } else if(this.guess(mid) == 1) {
                    i = mid + 1;
                } else {
                    j = mid;
                }
            }
            return i;
        }
    }

}

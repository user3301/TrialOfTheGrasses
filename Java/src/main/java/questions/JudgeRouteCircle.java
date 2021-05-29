package questions;

/**
 * Created by user3301 on 8/21/2017.
 * Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.

 The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.

 Example 1:
 Input: "UD"
 Output: true
 Example 2:
 Input: "LL"
 Output: false

 */
public class JudgeRouteCircle {

    public boolean judgeCircle(String moves) {
        int[] coordinate = {0,0};
        for (char letter : moves.toCharArray()) {
            switch (letter) {
                case 'U':
                    coordinate[1]++;
                    break;
                case 'D':
                    coordinate[1]--;
                    break;
                case 'L':
                    coordinate[0]++;
                    break;
                case 'R':
                    coordinate[0]--;
                    break;
            }
        }
        if(coordinate[0]==0 && coordinate[1]==0) {
            return true;
        }
        else {
            return false;
        }
    }

    public static void main(String[] args) {
        JudgeRouteCircle judgeRouteCircle = new JudgeRouteCircle();
        String str = "ULDR";
        Boolean result  = judgeRouteCircle.judgeCircle(str);
        System.out.println(result);
    }
}

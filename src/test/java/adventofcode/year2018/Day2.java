package adventofcode.year2018;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Day2 extends BaseTest {

    @Test
    public void runSilver() throws Exception {
        final ArrayList<String> data = readStringFromFile("year2018/day2/input.txt");
        int count2 = 0;
        int count3 = 0;
        for (final String input : data) {
            if (countAppear(input, 2)) {
                count2++;
            }
            if (countAppear(input, 3)) {
                count3++;
            }
        }
        System.out.println(count2 * count3);
    }

    boolean countAppear(String input, int count) {
        char[] chars = input.toCharArray();
        Map<Character, Integer> mapCharInt = new HashMap<>();
        for (char c : chars) {
            mapCharInt.compute(c, (k, v) -> v == null ? 1 : v + 1);
        }
        for (int i : mapCharInt.values()) {
            if (i == count) {
                return true;
            }
        }
        return false;
    }


    @Test
    public void runGold() throws Exception {
        final ArrayList<String> data = readStringFromFile("year2018/day2/input.txt");
        for (final String input : data) {
            for (final String input2 : data) {
                differCount(input, input2);

            }
        }
    }

    void differCount(String input, String input2) {
        char[] chars = input.toCharArray();
        char[] chars2 = input2.toCharArray();

        int differCount = 0;
        for (int i = 0; i < chars2.length; i++) {
            if (chars2[i] != chars[i]) {
                differCount++;
            }
        }

        if (differCount == 1) {
            System.out.println();
            for (int i = 0; i < chars2.length; i++) {
                if (chars2[i] == chars[i]) {
                    System.out.print(chars[i]);
                }
            }
        }
    }

}

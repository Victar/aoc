package adventofcode.year2018;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Day1 extends BaseTest {

    @Test
    public void runSilver() throws Exception {
        final ArrayList<String> data = readStringFromFile("year2018/day1/input.txt");

        long result = 0l;
        for (String current : data) {
            result += Long.parseLong(current);
        }
        System.out.println(result);
    }

    @Test
    public void runGold() throws Exception {

        final ArrayList<String> data = readStringFromFile("year2018/day1/input.txt");
        Set<Long> positions = new HashSet<>();
        boolean found = false;
        long result = 0l;
        while (!found) {
            for (String current : data) {
                result += Long.parseLong(current);
                if (positions.contains(result)) {
                    found = true;
                    System.out.println(result);
                    break;
                } else {
                    positions.add(result);
                }

            }
        }
    }
}

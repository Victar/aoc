package adventofcode.year2018;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

public class Day5 extends BaseTest {

    @Test
    public void runSilver() throws Exception {
        final ArrayList<String> data = readStringFromFile("year2018/day5/input.txt");
        char[] arr = data.get(0).toCharArray();
        List<Character> characters = new ArrayList<>();
        for (char curC : arr) {
            characters.add(curC);
        }
        System.out.println(getSum(characters));
    }

    @Test
    public void runGold() throws Exception {
        final ArrayList<String> data = readStringFromFile("year2018/day5/input.txt");
        char[] arr = data.get(0).toCharArray();
        List<Character> characters = new ArrayList<>();
        for (char curC : arr) {
            characters.add(curC);
        }
        int min = Integer.MAX_VALUE;
        for (char cur = 'a'; cur <= 'z'; cur++) {
            List<Character> charactersFiltered = new ArrayList<>(characters);
            charactersFiltered.removeAll(List.of(cur, (char) (cur - 32)));
            int curSum = getSum(charactersFiltered);
            if (curSum < min) {
                min = curSum;
            }
        }
        System.out.println(min);
    }

    public int getSum(List<Character> characters) {
        boolean doSearch = true;

        List<Character> charactersNew = new ArrayList<>();
        while (doSearch) {
            doSearch = false;
            for (int i = 0; i < characters.size(); i++) {
                Character cur = characters.get(i);

                Character next = i < characters.size() - 1 ? characters.get(i + 1) : 100;
                if (Math.abs(cur - next) == 32) {
                    i = i + 1;
                    doSearch = true;
                } else {
                    charactersNew.add(cur);
                }
            }

            characters = charactersNew;
            charactersNew = new ArrayList<>();
        }
        return characters.size();
    }


}

package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;

public class Day4 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day4/input.txt");
		int total = 0;
		for (final String input : data) {
			final String[] inputArr = input.split(",");
			final String[] s1 = inputArr[0].split("-");
			final String[] s2 = inputArr[1].split("-");
			int s1b = Integer.parseInt(s1[0]);
			int s1e = Integer.parseInt(s1[1]);
			int s2b = Integer.parseInt(s2[0]);
			int s2e = Integer.parseInt(s2[1]);
			if (s1b <= s2b && s1e >= s2e || s2b <= s1b && s2e >= s1e) {
				total++;
			}
		}
		System.out.println(total);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day4/input.txt");
		int total = 0;
		for (final String input : data) {
			final String[] inputArr = input.split(",");
			final String[] s1 = inputArr[0].split("-");
			final String[] s2 = inputArr[1].split("-");
			int s1b = Integer.parseInt(s1[0]);
			int s1e = Integer.parseInt(s1[1]);
			int s2b = Integer.parseInt(s2[0]);
			int s2e = Integer.parseInt(s2[1]);
			if (s1b <= s2b && s2b <= s1e || s2b <= s1b && s1b <= s2e) {
				total++;
			}
		}
		System.out.println(total);
	}

}

package adventofcode.year2022;

import adventofcode.BaseTest;
import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Day1 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day1/input.txt");
		int max = Integer.MIN_VALUE;
		int current = 0;
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				if (current > max) {
					max = current;
				}
				current = 0;
			} else {
				current += Integer.parseInt(input);
			}
		}
		System.out.println(max);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day1/input.txt");
		int current = 0;
		List<Integer> total = new ArrayList<>();
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				total.add(current);
				current = 0;
			} else {
				current += Integer.parseInt(input);
			}
		}
		Collections.sort(total);
		int top3 = new ArrayList<>(total.subList(total.size() - 3, total.size())).stream().mapToInt(Integer::intValue).sum();
		System.out.println(top3);

	}

}

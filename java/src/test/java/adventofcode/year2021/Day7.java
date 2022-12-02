package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day7 extends BaseTest {

	@Test public void runSilver() throws Exception {
		run(false);
	}

	@Test public void runGold() throws Exception {
		run(true);
	}

	void run(final boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day7/input.txt");
		final String[] arr = data.get(0).split(",");
		final ArrayList<Integer> pos = new ArrayList<Integer>();

		for (final String input : arr) {
			pos.add(Integer.parseInt(input));
		}
		long min = Long.MAX_VALUE;
		final int max = pos.stream().max(Integer::compareTo).get();

		for (int i = 0; i < max; i++) {
			long current = 0;
			for (int j = 0; j < pos.size(); j++) {
				final int size = Math.abs(i - pos.get(j));
				final int length = isGold ? (size * (size + 1)) / 2 : size;
				current += length;
			}
			min = Math.min(current, min);
		}
		System.out.println(min);
	}

}

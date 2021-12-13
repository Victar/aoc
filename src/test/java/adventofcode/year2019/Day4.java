package adventofcode.year2019;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day4 extends BaseTest {

	@Test public void runSilver() throws Exception {
		runAny(false);
	}

	@Test public void runGold() throws Exception {
		runAny(true);
	}

	public void runAny(final boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day4/input.txt");
		final String[] arr = data.get(0).split("-");
		final int start = Integer.parseInt(arr[0]);
		final int end = Integer.parseInt(arr[1]);
		int count = 0;
		for (int i = start; i <= end; i++) {
			boolean foundDouble = false;
			boolean valid = true;
			final String is = i + "";
			char prev = is.charAt(0);
			for (int j = 1; j < 6; j++) {
				final char current = is.charAt(j);
				valid = valid && prev <= current;
				if (isGold) {
					foundDouble = foundDouble || (prev == current)  // Check silver conditions
							&& (j + 1 >= 6 || is.charAt(j + 1) != current) // Check not triplet left
							&& (j - 2 < 0 || current != is.charAt(j - 2)); // Check not triplet right
				} else {
					foundDouble = foundDouble || current == prev;
				}
				prev = current;
			}
			if (foundDouble && valid) {
				count++;
			}
		}
		System.out.println(count);
	}

}

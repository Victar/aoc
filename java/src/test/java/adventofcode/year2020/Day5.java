package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Collections;

import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day5 extends BaseTest {

	@Test @Ignore public void singleCheck() {
		System.out.println(findSeat("FBFBBFFRLR"));
		System.out.println(findSeat("BBFFBBFRLL"));

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day5/input.txt");
		int count = 0;
		for (final String password : data) {
			final int current = findSeat(password);
			count = Math.max(count, current);
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day5/input.txt");
		final ArrayList<Integer> seats = new ArrayList();
		for (final String password : data) {
			seats.add(findSeat(password));
		}
		Collections.sort(seats);
		for (int i = 0; i < seats.size(); i++) {
			if (seats.get(i) != i + seats.get(0)) {
				System.out.println(seats.get(i - 1));
				return;
			}
		}
	}

	public int findSeat(final String input) {
		final String raw = input.substring(0, input.length() - 3);
		final String column = input.substring(input.length() - 3);
		final int rawInt = Integer.parseInt(StringUtils.replaceEach(raw, new String[] { "F", "B" }, new String[] { "0", "1" }), 2);
		final int columnInt = Integer.parseInt(StringUtils.replaceEach(column, new String[] { "R", "L" }, new String[] { "1", "0" }), 2);
		return rawInt * 8 + columnInt;
	}
}
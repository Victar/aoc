package adventofcode.year2020;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day15 extends BaseTest {

	public static final int[] data = { 18, 8, 0, 5, 4, 1, 20 };

	public static final int SIZE_SILVER = 2020;
	public static final int SIZE_GOLD = 30000000;

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final Map<Integer, Integer> map = new LinkedHashMap<Integer, Integer>();
		for (int i = 0; i < data.length; i++) {
			map.put(i, data[i]);
		}
		System.out.println(map);
		for (int i = data.length; i < SIZE_SILVER; i++) {
			final int number = map.get(i - 1);
			if (isFirstTimeSpoken(i - 1, number, map)) {
				map.put(i, 0);
			} else {
				final int turnLast = getNumberTurn(i - 1, number, map);
				final int turnBefore = getNumberTurn(i - 2, number, map);

				map.put(i, turnLast - turnBefore);
			}
			if (i % 1000 == 0) {
				System.out.println(i);
			}
		}
		System.out.println(map.get(SIZE_SILVER - 1));

	}

	@Test public void runGold() throws Exception {
		final Map<Integer, List<Integer>> map = new LinkedHashMap<>();
		int currentNumber = 0;
		for (int i = 0; i < data.length; i++) {
			currentNumber = data[i];
			map.putIfAbsent(currentNumber, new ArrayList<>());
			map.get(currentNumber).add(i + 1);
		}
		for (int i = data.length; i < SIZE_GOLD; i++) {
			if (map.get(currentNumber).size() == 1) {
				currentNumber = 0;
			} else {
				final List<Integer> l = map.get(currentNumber);
				currentNumber = l.get(l.size() - 1) - l.get(l.size() - 2);
			}
			map.putIfAbsent(currentNumber, new ArrayList<>());
			map.get(currentNumber).add(i + 1);
			if (i % 1000000 == 0) {
				System.out.println(i);
			}
		}
		System.out.println(currentNumber);
	}

	public int getNumberTurn(final int index, final int number, final Map<Integer, Integer> map) {
		for (int i = index; i >= 0; i--) {
			if (number == map.get(i)) {
				return i;
			}
		}
		System.out.println("Something wrong!");
		return -1;
	}

	public boolean isFirstTimeSpoken(final int index, final int number, final Map<Integer, Integer> map) {
		for (int i = 0; i < index; i++) {
			if (number == map.get(i)) {
				return false;
			}
		}
		return true;
	}

}


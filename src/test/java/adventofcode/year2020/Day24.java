package adventofcode.year2020;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day24 extends BaseTest {

	@Test public void singleCheck() {
		System.out.println(stringToHex("esew"));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day24/input.txt");
		final Map<Hex, Integer> map = new HashMap<>();
		for (final String input : data) {
			if (StringUtils.isNotEmpty(input)) {
				final Hex hex = stringToHex(input);
				if (map.containsKey(hex)) {
					map.put(hex, map.get(hex) + 1);
				} else {
					map.put(hex, 1);
				}
			}
		}
		int count = 0;
		for (final Map.Entry<Hex, Integer> entry : map.entrySet()) {
			System.out.println(entry.getKey() + " => " + entry.getValue());
			if (entry.getValue() % 2 == 1) {
				count++;
			}
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day24/input.txt");
		List<Hex> list = new ArrayList<>();
		for (final String input : data) {
			if (StringUtils.isNotEmpty(input)) {
				final Hex hex = stringToHex(input);
				if (list.contains(hex)) {
					list.remove(hex);
				} else {
					list.add(hex);
				}
			}
		}
		int SIZE = 50;
		for (int i = 0; i < 100; i++) {
			list = makeRound(list, ++SIZE);
		}
		System.out.println(list.size());
	}

	public List<Hex> makeRound(final List<Hex> currentList, final int size) {
		final List<Hex> result = new ArrayList<>();
		for (int i = -size; i <= size; i++) {
			for (int j = -size; j <= size; j++) {
				final Hex current = new Hex(i, j);
				final int countBlackNeighbors = countBlackNeighbors(current, currentList);
				final boolean isBlack = currentList.contains(current);
				if (isBlack && countBlackNeighbors != 0 && countBlackNeighbors <= 2) {
					result.add(current);
				}
				if (!isBlack && countBlackNeighbors == 2) {
					result.add(current);
				}
			}
		}
		return result;
	}

	public int countBlackNeighbors(final Hex hex, final List<Hex> currentList) {
		int count = 0;
		if (currentList.contains(new Hex(hex.x + 1, hex.y))) {
			count++;
		}
		if (currentList.contains(new Hex(hex.x + 1, hex.y - 1))) {
			count++;
		}
		if (currentList.contains(new Hex(hex.x, hex.y - 1))) {
			count++;
		}
		if (currentList.contains(new Hex(hex.x - 1, hex.y))) {
			count++;
		}
		if (currentList.contains(new Hex(hex.x - 1, hex.y + 1))) {
			count++;
		}
		if (currentList.contains(new Hex(hex.x, hex.y + 1))) {
			count++;
		}

		return count;
	}

	public Hex stringToHex(String input) {
		int x = 0;
		int y = 0;
		while (StringUtils.isNotBlank(input)) {
			if (input.startsWith("w")) {
				input = input.substring(1);
				x = x - 1;
			}
			if (input.startsWith("e")) {
				input = input.substring(1);
				x = x + 1;
			}
			if (input.startsWith("sw")) {
				input = input.substring(2);
				y = y - 1;
			}
			if (input.startsWith("se")) {
				input = input.substring(2);
				y = y - 1;
				x = x + 1;
			}
			if (input.startsWith("nw")) {
				input = input.substring(2);
				x = x - 1;
				y = y + 1;
			}
			if (input.startsWith("ne")) {
				input = input.substring(2);
				y = y + 1;
			}
		}
		return new Hex(x, y);
	}

	@Data public static class Hex {

		int x;
		int y;

		public Hex(final int x, final int y) {
			this.x = x;
			this.y = y;
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Hex hex = (Hex) o;

			if (this.x != hex.x) return false;
			return this.y == hex.y;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			return result;
		}

		@Override public String toString() {
			return "Hex{" + "x=" + this.x + ", y=" + this.y + '}';
		}
	}
}


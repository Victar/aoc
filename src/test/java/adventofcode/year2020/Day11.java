package adventofcode.year2020;

import java.util.ArrayList;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day11 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day11/input.txt");
		printArray(data);
		int count = -1;
		while (count != countSeats(data)) {
			count = countSeats(data);
			doRound(data);
		}
		System.out.println("Count:  " + countSeats(data));

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day11/input.txt");
		//		org.junit.Assert.assertEquals(3, hasAdjacentGold( data,  0, 9, '#'));
		//		org.junit.Assert.assertEquals(7, hasAdjacentGold( data,  1, 1, '#'));
		//		org.junit.Assert.assertEquals(5, hasAdjacentGold( data,  1, 9, '#'));
		printArray(data);
		System.out.println("------------------  " + countSeats(data));
		int count = -1;
		while (count != countSeats(data)) {
			count = countSeats(data);
			System.out.println("------------------  ");
			printArray(data);
			doRound(data);

		}
		System.out.println("Count:  " + countSeats(data));

	}

	public int countSeats(final ArrayList<String> data) {
		int total = 0;
		for (final String s : data) {
			total += StringUtils.countMatches(s, '#');
		}
		return total;
	}

	public void printArray(final ArrayList<String> data) {
		for (final String s : data) {
			System.out.println(s);
		}
	}

	public void doRound(final ArrayList<String> data) {
		final ArrayList<String> dataClone = (ArrayList<String>) data.clone();
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.get(i).length(); j++) {
				final char c = data.get(i).charAt(j);
				if (c == '.') {
					//do nothing
				}
				if (c == 'L') {
					if (hasAdjacentGold(dataClone, i, j, '#') == 0) {
						final StringBuilder currentString = new StringBuilder(data.get(i));
						currentString.setCharAt(j, '#');
						data.set(i, currentString.toString());
					}
				}
				if (c == '#') {
					if (hasAdjacentGold(dataClone, i, j, '#') >= 5) {
						final StringBuilder currentString = new StringBuilder(data.get(i));
						currentString.setCharAt(j, 'L');
						data.set(i, currentString.toString());
					}
				}

			}
		}
	}

	public int hasAdjacentGold(final ArrayList<String> data, final int i, final int j, final char toCheck) {
		int count = 0;
		//Check Row lest X
		final String currentString = data.get(i);
		final int currentLength = currentString.length();

		final String currentStringLeft = currentString.substring(0, j);
		for (int l = currentStringLeft.length() - 1; l >= 0; l--) {
			if (currentStringLeft.charAt(l) == 'L') {
				break;
			}
			if (currentStringLeft.charAt(l) == toCheck) {
				count++;
				break;
			}
		}

		final String currentStringRight = currentString.substring(j + 1);
		for (int l = 0; l < currentStringRight.length(); l++) {
			if (currentStringRight.charAt(l) == 'L') {
				break;
			}
			if (currentStringRight.charAt(l) == toCheck) {
				count++;
				break;
			}
		}

		for (int k = i - 1; k >= 0; k--) {
			if (data.get(k).charAt(j) == 'L') {
				break;
			}
			if (data.get(k).charAt(j) == toCheck) {
				count++;
				break;
			}
		}
		for (int k = i + 1; k < data.size(); k++) {
			if (data.get(k).charAt(j) == 'L') {
				break;
			}
			if (data.get(k).charAt(j) == toCheck) {
				count++;
				break;
			}
		}
		for (int k = i + 1, l = j + 1; k < data.size() && l < currentLength; k++, l++) {
			if (data.get(k).charAt(l) == 'L') {
				break;
			}

			if (data.get(k).charAt(l) == toCheck) {
				count++;
				break;
			}
		}

		for (int k = i + 1, l = j - 1; k < data.size() && l >= 0; k++, l--) {
			if (data.get(k).charAt(l) == 'L') {
				break;
			}
			if (data.get(k).charAt(l) == toCheck) {
				count++;
				break;
			}
		}

		for (int k = i - 1, l = j - 1; k >= 0 && l >= 0; k--, l--) {
			if (data.get(k).charAt(l) == 'L') {
				break;
			}

			if (data.get(k).charAt(l) == toCheck) {
				count++;
				break;
			}
		}

		for (int k = i - 1, l = j + 1; k >= 0 && l < currentLength; k--, l++) {
			if (data.get(k).charAt(l) == 'L') {
				break;
			}
			if (data.get(k).charAt(l) == toCheck) {
				count++;
				break;
			}
		}
		return count;
	}

	public int hasAdjacent(final ArrayList<String> data, final int i, final int j, final char toCheck) {
		int count = 0;

		try {
			if (data.get(i).charAt(j + 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i).charAt(j - 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i - 1).charAt(j) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i - 1).charAt(j - 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i - 1).charAt(j + 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i + 1).charAt(j - 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i + 1).charAt(j) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		try {
			if (data.get(i + 1).charAt(j + 1) == toCheck) {
				count++;
			}
		} catch (final Exception ex) {
		}
		return count;

	}

}


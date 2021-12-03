package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day20 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day20/input.txt");
		final List<Tile> tiles = new ArrayList<>();
		Tile currentTile = new Tile();
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				tiles.add(currentTile);
				currentTile = new Tile();
			} else {
				final Pattern p = Pattern.compile("\\d+");
				final Matcher m = p.matcher(input);
				if (m.find()) {
					currentTile.setId(Integer.parseInt(m.group()));
				} else {
					currentTile.addDataLine(input);
				}

			}
		}
		tiles.add(currentTile);
		for (final Tile tile : tiles) {
			tile.initTiles();
			tile.printTile();
		}
		final Map<Integer, Integer> tileBorderCount = new HashMap<>();
		for (final Tile tile : tiles) {
			for (final int i : tile.borderAll) {
				final Integer mapI = tileBorderCount.get(i);
				if (mapI != null) {
					tileBorderCount.put(i, mapI + 1);
				} else {
					tileBorderCount.put(i, 1);
				}
			}
		}
		final List<Integer> singleBorder = new ArrayList<>();
		for (final Map.Entry<Integer, Integer> entry : tileBorderCount.entrySet()) {
			if (entry.getValue() == 1) {
				singleBorder.add(entry.getKey());
			}
		}
		long result = 1;
		for (final Tile tile : tiles) {
			final int sbCount = tile.singleBorderCount(singleBorder);
			if (sbCount > 2) {
				System.out.println("sbCount" + sbCount);
				result *= tile.id;
				tile.printTile();
			}
		}
		System.out.println(tileBorderCount);
		System.out.println("Result: " + result);

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day20/input.txt");
		final List<Tile> tiles = new ArrayList<>();
		Tile currentTile = new Tile();
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				tiles.add(currentTile);
				currentTile = new Tile();
			} else {
				final Pattern p = Pattern.compile("\\d+");
				final Matcher m = p.matcher(input);
				if (m.find()) {
					currentTile.setId(Integer.parseInt(m.group()));
				} else {
					currentTile.addDataLine(input);
				}

			}
		}
		tiles.add(currentTile);
		for (final Tile tile : tiles) {
			tile.initTiles();
		}
		final Map<Integer, Integer> tileBorderCount = new HashMap<>();
		for (final Tile tile : tiles) {
			for (final int i : tile.borderAll) {
				final Integer mapI = tileBorderCount.get(i);
				if (mapI != null) {
					tileBorderCount.put(i, mapI + 1);
				} else {
					tileBorderCount.put(i, 1);
				}
			}
		}
		final List<Integer> singleBorder = new ArrayList<>();
		for (final Map.Entry<Integer, Integer> entry : tileBorderCount.entrySet()) {
			if (entry.getValue() == 1) {
				singleBorder.add(entry.getKey());
			}
		}
		int totalCount = 0;
		for (final Tile tile : tiles) {
			totalCount += tile.squareBorderCount(singleBorder);

		}
		final int MONSTER_GUESS = 40;

		for (int i = 0; i < MONSTER_GUESS && i < 10; i++) {
			System.out.println(totalCount - 15 * (MONSTER_GUESS + i));
			//			System.out.println(totalCount - 15*(MONSTER_GUESS-i));

		}
	}

	@Data class Tile {

		int id;
		int[] border;
		int[] borderFlipped;
		int[] borderAll;

		ArrayList<String> data = new ArrayList<String>();

		public void addDataLine(final String dataLine) {
			this.data.add(dataLine);
		}

		public void initTiles() {
			initBorders();
		}

		public int squareBorderCount(final List<Integer> singleBorder) {
			int count = 0;
			for (int i = 1; i < this.data.size() - 1; i++) {
				final String substring = this.data.get(i).substring(1, this.data.get(i).length() - 1);
				//				System.out.println(substring);
				count += StringUtils.countMatches(substring, "#");
			}
			final int singleBorderCount = singleBorderCount(singleBorder);
			for (int i = 0; i < singleBorder.size(); i++) {
				if (contains(this.border, singleBorder.get(i))) {
					final int borderCount = singleBorder.get(i);
					if (singleBorderCount > 2) {
						//						System.out.println(borderCount + " ->" +  Integer.toBinaryString(borderCount ));
					}
				}
			}
			return count;
		}

		public int singleBorderCount(final List<Integer> border) {
			int count = 0;
			for (final Integer b : border) {
				if (containsBorder(b)) {
					count++;
				}
			}
			return count;
		}

		public boolean containsBorder(final int border) {
			return contains(this.borderAll, border);
		}

		public void printTile() {
			System.out.println("Tile " + this.id + ":");
			for (final String s : this.data) {
				System.out.println(s);
			}
			System.out.println(Arrays.toString(this.border));
			System.out.println(Arrays.toString(this.borderFlipped));
			System.out.println(Arrays.toString(this.borderAll));
			System.out.println();
		}

		private boolean contains(final int[] arr, final int key) {
			return Arrays.stream(arr).anyMatch(i -> i == key);
		}

		public void initBorders() {
			if (this.border == null) {
				final StringBuilder sbX2 = new StringBuilder();
				final StringBuilder sbX4 = new StringBuilder();
				for (int i = 0; i < this.data.size(); i++) {
					sbX4.append(this.data.get(i).charAt(0));
					sbX2.append(this.data.get(i).charAt(this.data.size() - 1));
				}
				final String s1 = StringUtils.replaceEach(this.data.get(0), new String[] { ".", "#" }, new String[] { "0", "1" });
				final String s3 = StringUtils.replaceEach(StringUtils.reverse(this.data.get(this.data.size() - 1)), new String[] { ".", "#" },
						new String[] { "0", "1" });
				final String s2 = StringUtils.replaceEach(sbX2.toString(), new String[] { ".", "#" }, new String[] { "0", "1" });
				final String s4 = StringUtils.replaceEach(sbX4.reverse().toString(), new String[] { ".", "#" }, new String[] { "0", "1" });

				final int x1 = Integer.parseInt(s1, 2);
				final int x3 = Integer.parseInt(s3, 2);
				final int x2 = Integer.parseInt(s2, 2);
				final int x4 = Integer.parseInt(s4, 2);
				this.border = new int[] { x1, x2, x3, x4 };
				final int xf1 = Integer.parseInt(StringUtils.reverse(s1), 2);
				final int xf3 = Integer.parseInt(StringUtils.reverse(s3), 2);
				final int xf2 = Integer.parseInt(StringUtils.reverse(s2), 2);
				final int xf4 = Integer.parseInt(StringUtils.reverse(s4), 2);

				this.borderFlipped = new int[] { xf1, xf2, xf3, xf4 };
				this.borderAll = new int[] { x1, x2, x3, x4, xf1, xf2, xf3, xf4 };
			}
		}
	}

}


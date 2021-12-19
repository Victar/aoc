package adventofcode.year2021;

import java.util.*;
import java.util.stream.Collectors;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day19 extends BaseTest {

	@Test public void runBoth() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day19/input.txt");
		String currentTitle = StringUtils.EMPTY;
		List<Coord> currentCoords = new ArrayList<>();
		final List<Scanner> scanners = new ArrayList<>();
		boolean first = true;
		final List<Scanner> alligned = new ArrayList<>();
		for (final String input : data) {
			if (StringUtils.isNotEmpty(input)) {
				if (input.contains("scanner")) {
					if (StringUtils.isNotEmpty(currentTitle)) {
						final Scanner scanner = new Scanner(currentCoords, currentTitle);
						if (first) {
							scanner.setPosition(new Coord(0, 0, 0));
							alligned.add(scanner);
							scanner.setCordsAligned(scanner.getCords());
							first = false;
						}
						scanners.add(scanner);
					}
					currentTitle = input;
					currentCoords = new ArrayList<>();
				} else {
					final String[] arr = StringUtils.split(input, ",");
					final Coord currrent = new Coord(Integer.parseInt(arr[0]), Integer.parseInt(arr[1]), Integer.parseInt(arr[2]));
					currentCoords.add(currrent);
				}
			}
		}
		final Scanner scanner = new Scanner(currentCoords, currentTitle);
		scanners.add(scanner);
		boolean foundAll = false;
		while (!foundAll) {
			final int alignedSize = alligned.size();
			for (int i = 0; i < alignedSize; i++) {
				for (final Scanner toCheck : scanners) {
					if (!alligned.contains(toCheck)) {
						final boolean matched = alligned.get(i).findMath(toCheck);
						if (matched) {
							alligned.add(toCheck);
						}
					}
				}
			}
			foundAll = alligned.size() == scanners.size();
		}
		final Set<Coord> beacons = new HashSet<>();
		for (final Scanner current : scanners) {
			beacons.addAll(current.getAligned());
		}
		int maxDistance = Integer.MIN_VALUE;
		for (int i = 0; i < scanners.size(); i++) {
			for (int j = i; j < scanners.size(); j++) {
				maxDistance = Math.max(maxDistance, scanners.get(i).countDistance(scanners.get(j)));
			}
		}
		System.out.println(beacons.size());
		System.out.println(maxDistance);
	}

	@Data class Scanner {

		String title;
		List<Coord> cords;
		Map<Integer, List<Coord>> positionCords = new HashMap<>();
		//Set once aligned;
		List<Coord> cordsAligned;
		Coord position;
		Scanner alignedWith;

		public Scanner(final List<Coord> cords, final String title) {
			this.title = title;
			this.cords = cords;
			// It will generate 48 positions some of them might be duplicated or not valid :)
			// Duplicated should be fine
			final boolean[] xDirection = { true, false };
			final boolean[] yDirection = { true, false };
			final boolean[] zDirection = { true, false };
			for (int x = 0; x < xDirection.length; x++) {
				for (int y = 0; y < yDirection.length; y++) {
					for (int z = 0; z < zDirection.length; z++) {
						generateScanner(xDirection[x], yDirection[y], zDirection[z], cords);
					}
				}
			}
		}

		private void generateScanner(final boolean isXdirect, final boolean isYdirect, final boolean isZdirect, final List<Coord> cords) {
			final List<Coord> xyz = new ArrayList<>();
			final List<Coord> xzy = new ArrayList<>();
			final List<Coord> yxz = new ArrayList<>();
			final List<Coord> yzx = new ArrayList<>();
			final List<Coord> zxy = new ArrayList<>();
			final List<Coord> zyx = new ArrayList<>();
			final int x = isXdirect ? 1 : -1;
			final int y = isYdirect ? 1 : -1;
			final int z = isZdirect ? 1 : -1;

			for (final Coord cord : cords) {
				xyz.add(new Coord(x * cord.x, y * cord.y, z * cord.z));
				xzy.add(new Coord(x * cord.x, y * cord.z, z * cord.y));
				yxz.add(new Coord(x * cord.y, y * cord.x, z * cord.z));
				yzx.add(new Coord(x * cord.y, y * cord.z, z * cord.x));
				zxy.add(new Coord(x * cord.z, y * cord.x, z * cord.y));
				zyx.add(new Coord(x * cord.z, y * cord.y, z * cord.x));

			}
			this.positionCords.put(this.positionCords.size(), xyz);
			this.positionCords.put(this.positionCords.size(), xzy);
			this.positionCords.put(this.positionCords.size(), yxz);
			this.positionCords.put(this.positionCords.size(), yzx);
			this.positionCords.put(this.positionCords.size(), zxy);
			this.positionCords.put(this.positionCords.size(), zyx);
		}

		public int countDistance(final Scanner another) {
			return Math.abs(this.position.x - another.position.x) + Math.abs(this.position.y - another.position.y) + Math.abs(
					this.position.z - another.position.z);
		}

		public boolean findMath(final Scanner scannerToCheck) {
			for (final Map.Entry<Integer, List<Coord>> entry : scannerToCheck.positionCords.entrySet()) {
				final List<Coord> scannerToCheckCords = entry.getValue();
				final Map<Coord, Long> matches =
						this.cordsAligned.stream().map(c -> scannerToCheckCords.stream().map(o -> o.sub(c)).collect(Collectors.toList()))
								.flatMap(Collection::stream).collect(Collectors.groupingBy(t -> t, Collectors.counting()));
				for (final Map.Entry<Coord, Long> match : matches.entrySet()) {
					if (match.getValue() >= 12) {
						scannerToCheck.position = this.position.sub(match.getKey());
						scannerToCheck.cordsAligned = scannerToCheckCords;
						scannerToCheck.alignedWith = this;
						//						System.out.println("found! allignedMap: " + entry.getKey() + "  " + match + "this: " + this.title + " aligned: "
						//								+ scannerToCheck.title);
						return true;
					}
				}
			}
			return false;
		}

		public Set<Coord> getAligned() {
			final Set<Coord> result = new HashSet<>();
			for (final Coord cord : this.cordsAligned) {
				result.add(cord.add(this.position));
			}
			return result;
		}

		public void printAligned() {
			System.out.println(this.title + " " + this.position);
			for (final Coord cord : this.cordsAligned) {
				System.out.println("Cord " + cord + " " + cord.sub(this.position));
			}
		}

	}

	@Data class Coord {

		int x;
		int y;
		int z;

		public Coord(final int x, final int y, final int z) {
			this.x = x;
			this.y = y;
			this.z = z;
		}

		public Coord add(final Coord other) {
			return new Coord(this.x + other.x, this.y + other.y, this.z + other.z);
		}

		public Coord sub(final Coord other) {
			return new Coord(this.x - other.x, this.y - other.y, this.z - other.z);
		}

		@Override public String toString() {
			return "(" + this.x + "," + this.y + "," + this.z + ")";
		}
	}
}
